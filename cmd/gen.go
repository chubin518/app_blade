/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"app_blade/pkg/config"
	"app_blade/pkg/database"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	"datetime": func(columnType gorm.ColumnType) (dataType string) { return "database.LocalTime" },
}

// generatorCmd represents the generator command
var generatorCmd = &cobra.Command{
	Use:   "gen",
	Short: "grom code generation",
	Run: func(cmd *cobra.Command, args []string) {
		dsn := config.Default().GetString("database.datasource.product.dsn")
		db, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic(err)
		}
		g := gen.NewGenerator(gen.Config{
			ModelPkgPath: "internal/model",
		})
		// g.WithDataTypeMap(dataMap)
		g.WithJSONTagNameStrategy(func(columnName string) string {
			if columnName == "password" {
				return "-"
			}
			return columnName
		})
		g.UseDB(db)
		g.WithDataTypeMap(dataMap)
		g.GenerateAllTable()
		g.Execute()
	},
}

func init() {
	rootCmd.AddCommand(generatorCmd)
}

type CommonMethod struct {
	ID       int64
	CreateAt database.LocalTime
	CreateBy string
	UpdateAt database.LocalTime
	UpdateBy string
}

func (cm *CommonMethod) BeforeCreate(tx *gorm.DB) (err error) {
	if cm.CreateAt.IsZero() {
		cm.CreateAt = database.NowTime()
	}
	if cm.UpdateAt.IsZero() {
		cm.UpdateAt = database.NowTime()
	}
	return nil
}

func (cm *CommonMethod) BeforeUpdate(tx *gorm.DB) (err error) {
	if cm.UpdateAt.IsZero() {
		cm.UpdateAt = database.NowTime()
	}
	return nil
}
