package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Food interface {
	GetFood(ctx context.Context, id string) (service_models.Food, error)
	GetFoods(ctx context.Context) ([]service_models.Food, error)
	CreateFood(ctx context.Context, food service_models.Food) error
	UpdateFood(ctx context.Context, id string) error
}

type foodRepo struct {
	collection *mongo.Collection
}

func (f *foodRepo) GetFoods(ctx context.Context) ([]service_models.Food, error) {
	return nil, nil
}

func (f *foodRepo) GetFood(ctx context.Context, id string) (service_models.Food, error) {
	return service_models.Food{}, nil
}

func (f *foodRepo) CreateFood(ctx context.Context, food service_models.Food) error {
	return nil
}

func (f *foodRepo) UpdateFood(ctx context.Context, id string) error {
	return nil
}

func NewFoodRepo(db *mongo.Database) Food {
	return &foodRepo{
		collection: db.Collection("food"),
	}
}
