package user

import (
	"context"
	"simple-crud/internal/controller/user/data"
	"simple-crud/internal/controller/user/repository"
)

type DefaultService struct {
	repo repository.Repository
}

func NewDefaultService(repo repository.Repository) *DefaultService {
	return &DefaultService{repo}
}

func (s *DefaultService) Create(ctx context.Context, username, email, password string) error {
	return s.repo.Create(ctx, username, email, password)
}

func (s *DefaultService) Read(ctx context.Context, id int64) (*data.User, error) {
	return s.repo.Read(ctx, id)
}

func (s *DefaultService) Update(ctx context.Context, username, email, password string) error {
	return s.repo.Update(ctx, username, email, password)
}

func (s *DefaultService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
