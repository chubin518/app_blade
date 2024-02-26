package service

import (
	"app_blade/internal/model"
	"app_blade/internal/repository"
	"app_blade/pkg/logging"
	"context"
)

var _ BladeProductService = (*DefaultBladeProductService)(nil)

func NewBladeProductService(repository repository.BladeProductRepository, logger logging.Provider) BladeProductService {
	return &DefaultBladeProductService{
		repository: repository,
		logger:     logger,
	}
}

type BladeProductService interface {
	List(ctx context.Context) ([]*model.BladeProduct, error)
	SaveOrUpdate(ctx context.Context, m *model.BladeProduct) error
}

type DefaultBladeProductService struct {
	logger     logging.Provider
	repository repository.BladeProductRepository
}

// Save implements BladeProductService.
func (s *DefaultBladeProductService) SaveOrUpdate(ctx context.Context, m *model.BladeProduct) error {
	return s.repository.SaveOrUpdate(ctx, m)
}

// List implements BladeProductService.
func (s *DefaultBladeProductService) List(ctx context.Context) ([]*model.BladeProduct, error) {
	return s.repository.List(ctx)
}
