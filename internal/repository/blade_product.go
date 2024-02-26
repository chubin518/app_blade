package repository

import (
	"app_blade/internal/model"
	"app_blade/pkg/logging"
	"context"

	"gorm.io/gorm"
)

var _ BladeProductRepository = (*MySQLBladeProductRepository)(nil)

func NewBladeProductRepository(db *gorm.DB, logger logging.Provider) BladeProductRepository {
	return &MySQLBladeProductRepository{
		db:     db,
		logger: logger,
	}
}

type BladeProductRepository interface {
	List(ctx context.Context) ([]*model.BladeProduct, error)
	SaveOrUpdate(ctx context.Context, m *model.BladeProduct) error
}

type MySQLBladeProductRepository struct {
	logger logging.Provider
	db     *gorm.DB
}

// Save implements BladeProductRepository.
func (r *MySQLBladeProductRepository) SaveOrUpdate(ctx context.Context, m *model.BladeProduct) error {
	return r.db.WithContext(ctx).Save(m).Error
}

// List implements BladeProductRepository.
func (r *MySQLBladeProductRepository) List(ctx context.Context) (l []*model.BladeProduct, err error) {
	err = r.db.WithContext(ctx).Find(&l).Error
	return l, err
}
