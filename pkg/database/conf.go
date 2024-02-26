package database

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Config struct {
	Primary      string                `mapstructure:"primary"`
	MaxIdleTime  time.Duration         `mapstructure:"max_idle_time"`
	MaxLifeTime  time.Duration         `mapstructure:"max_life_time"`
	MaxIdleConns int                   `mapstructure:"max_idle_conns"`
	MaxOpenConns int                   `mapstructure:"max_open_conns"`
	DataSource   map[string]DataSource `mapstructure:"datasource"`
}

func (cfg *Config) dialector() gorm.Dialector {
	var dsn string
	for k, ds := range cfg.DataSource {
		if len(cfg.Primary) == 0 {
			dsn = ds.DSN
			break
		} else if cfg.Primary == k {
			dsn = ds.DSN
			break
		}
	}
	return mysql.Open(dsn)
}

func (cfg *Config) dbresolver() *dbresolver.DBResolver {
	var resolver *dbresolver.DBResolver
	for _, ds := range cfg.DataSource {
		if resolver == nil {
			resolver = dbresolver.Register(ds.dbresolver(), ds.Tables...)
		} else {
			resolver.Register(ds.dbresolver(), ds.Tables...)
		}
	}
	return resolver.
		SetConnMaxIdleTime(cfg.MaxIdleTime).
		SetConnMaxLifetime(cfg.MaxLifeTime).
		SetMaxIdleConns(cfg.MaxIdleConns).
		SetMaxOpenConns(cfg.MaxOpenConns)
}

type DataSource struct {
	// 主库dsn
	DSN string `mapstructure:"dsn"`
	// 从库dsn
	Replicas []string `mapstructure:"replicas"`
	// 工作表
	Tables []any `mapstructure:"tables"`
}

func (ds *DataSource) dbresolver() dbresolver.Config {
	replicas := make([]gorm.Dialector, 0)
	if len(ds.Replicas) >= 1 {
		for _, replica := range ds.Replicas {
			replicas = append(replicas, mysql.Open(replica))
		}
	}
	return dbresolver.Config{
		Sources:           []gorm.Dialector{mysql.Open(ds.DSN)},
		Replicas:          replicas,
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Primary:      "",
		MaxIdleConns: 4,
		MaxOpenConns: 8,
		MaxIdleTime:  10 * time.Second,
		MaxLifeTime:  30 * time.Second,
		DataSource:   make(map[string]DataSource),
	}
}
