package service

import (
	"bluenote/internal/domain"
	"bluenote/internal/repository"
	"context"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(ctx context.Context, u domain.User) error {
	err := s.repo.Create(ctx, u)
	if err != nil {
		return err
	}
	return nil
}
