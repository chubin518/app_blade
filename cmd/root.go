/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"app_blade/pkg/config"
	"app_blade/pkg/logging"
	"os"

	"github.com/spf13/cobra"
)

var env string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "app_blade",
	Short: "description of app blade",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := CreateApp(config.WithName(env))
		if err != nil {
			logging.Default().Errorf("app init error: %s", err)
			return
		}
		if err := app.Run(); err != nil {
			logging.Default().Errorf("app run error: %s", err)
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&env, "env", "e", "dev", "设置运行环境")
}
