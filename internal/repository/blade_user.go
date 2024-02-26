package repository

import (
	"app_blade/internal/model"
	"app_blade/pkg/logging"
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ BladeUserRepository = (*MySQLBladeUserRepository)(nil)

func NewBladeUserRepository(db *gorm.DB, logger logging.Provider) BladeUserRepository {
	return &MySQLBladeUserRepository{
		logger: logger,
		db:     db,
	}
}

type BladeUserRepository interface {
	Get(ctx context.Context, id int) (*model.BladeUser, error)
	SaveOrUpdate(ctx context.Context, m *model.BladeUser) error
}

type MySQLBladeUserRepository struct {
	logger logging.Provider
	db     *gorm.DB
}

// Save implements BladeUserRepository.
func (r *MySQLBladeUserRepository) SaveOrUpdate(ctx context.Context, m *model.BladeUser) error {
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"update_at",
			"age",
			"create_by",
			"update_by",
			"password",
			"email",
			"address",
			"name",
		}),
	}).Create(m).Error
}

// Get implements BladeUserRepository.
func (r *MySQLBladeUserRepository) Get(ctx context.Context, id int) (m *model.BladeUser, err error) {
	err = r.db.WithContext(ctx).First(&m, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		m = nil
	}
	return m, err
}
