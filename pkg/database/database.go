package database

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewLogger, NewOptions, New)

func New(logger *Logger, options ...Option) (*gorm.DB, error) {
	cfg := DefaultConfig()
	for _, apply := range options {
		apply(cfg)
	}

	db, err := gorm.Open(cfg.dialector(), &gorm.Config{Logger: logger})
	if err != nil {
		return nil, err
	}

	if err := db.Use(cfg.dbresolver()); err != nil {
		return nil, err
	}

	return db, nil
}
