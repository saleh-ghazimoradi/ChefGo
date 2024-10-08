package service

import (
	"context"
	"github.com/saleh-ghazimoradi/ChefGo/internal/repository"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
)

type User interface {
	GetUser(ctx context.Context, id string) (service_models.User, error)
	GetUsers(ctx context.Context) ([]service_models.User, error)
	CreateUser(ctx context.Context, menu service_models.User) error
	UpdateUser(ctx context.Context, id string) error
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, password string) (bool, error)
}

type userService struct {
	userRepo repository.User
}

func (s *userService) GetUser(ctx context.Context, id string) (service_models.User, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *userService) GetUsers(ctx context.Context) ([]service_models.User, error) {
	return s.userRepo.GetUsers(ctx)
}

func (s *userService) CreateUser(ctx context.Context, user service_models.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, id string) error {
	return s.userRepo.UpdateUser(ctx, id)
}

func (s *userService) HashPassword(password string) (string, error) {
	return "", nil
}

func (s *userService) VerifyPassword(hashedPassword, password string) (bool, error) {
	return false, nil
}

func NewUserService(userRepo repository.User) User {
	return &userService{userRepo: userRepo}
}
