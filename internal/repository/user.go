package repository

import (
	"bluenote/internal/domain"
	"bluenote/internal/repository/dao"
	"context"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (repo *UserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (repo *UserRepository) FindByEmail(ctx context.Context, u domain.User) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, u)
	return u, err
}

func (repo *UserRepository) FindById(ctx context.Context, u domain.User) (domain.User, error) {
	u, err := repo.dao.FindById(ctx, u)
	return u, err
}

func (repo *UserRepository) EditById(ctx context.Context, u domain.User) error {
	err := repo.dao.EditById(ctx, u)
	if err != nil {
		return err
	}
	return nil
}
