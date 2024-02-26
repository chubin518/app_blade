package service

import (
	"app_blade/internal/model"
	"app_blade/internal/repository"
	"app_blade/pkg/logging"
	"context"
)

var _ BladeUserService = (*DefaultBladeUserService)(nil)

func NewBladeUserService(repository repository.BladeUserRepository, logger logging.Provider) BladeUserService {
	return &DefaultBladeUserService{
		repository: repository,
		logger:     logger,
	}
}

type BladeUserService interface {
	Get(ctx context.Context, id int) (*model.BladeUser, error)
	SaveOrUpdate(ctx context.Context, m *model.BladeUser) error
}

type DefaultBladeUserService struct {
	logger     logging.Provider
	repository repository.BladeUserRepository
}

// Save implements BladeUserService.
func (s *DefaultBladeUserService) SaveOrUpdate(ctx context.Context, m *model.BladeUser) error {
	return s.repository.SaveOrUpdate(ctx, m)
}

func (s *DefaultBladeUserService) Get(ctx context.Context, id int) (*model.BladeUser, error) {
	return s.repository.Get(ctx, id)
}
