package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	GetUser(ctx context.Context, id string) (service_models.User, error)
	GetUsers(ctx context.Context) ([]service_models.User, error)
	CreateUser(ctx context.Context, menu service_models.User) error
	UpdateUser(ctx context.Context, id string) error
}

type userRepo struct {
	collection *mongo.Collection
}

func (u *userRepo) GetUser(ctx context.Context, id string) (service_models.User, error) {
	return service_models.User{}, nil
}

func (u *userRepo) GetUsers(ctx context.Context) ([]service_models.User, error) {
	return nil, nil
}

func (u *userRepo) CreateUser(ctx context.Context, menu service_models.User) error {
	return nil
}

func (u *userRepo) UpdateUser(ctx context.Context, id string) error {
	return nil
}

func NewUserRepo(db *mongo.Database) User {
	return &userRepo{
		collection: db.Collection("user"),
	}
}
