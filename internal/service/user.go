package service

import (
	"bluenote/internal/domain"
	"bluenote/internal/repository"
	"bluenote/internal/repository/dao"
	"context"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail    = dao.ErrUserDuplicateEmail
	ErrInvalidUserOrPassword = dao.ErrInvalidUserOrPassword
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(ctx context.Context, u domain.User) error {
	pwdEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(pwdEncrypted)
	err = s.repo.Create(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(ctx context.Context, u domain.User) (domain.User, error) {
	user, err := s.repo.FindByEmail(ctx, u)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return user, ErrInvalidUserOrPassword
	}
	return user, nil
}

func (s *UserService) Profile(ctx context.Context, u domain.User) (domain.User, error) {
	user, err := s.repo.FindById(ctx, u)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) Edit(ctx context.Context, u domain.User) error {
	err := s.repo.EditById(ctx, u)
	return err
}
