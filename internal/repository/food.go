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

type FoodRepo struct {
	collection *mongo.Collection
}

func (f *FoodRepo) GetFoods(ctx context.Context) ([]service_models.Food, error) {
	return nil, nil
}

func (f *FoodRepo) GetFood(ctx context.Context, id string) (service_models.Food, error) {
	return service_models.Food{}, nil
}

func (f *FoodRepo) CreateFood(ctx context.Context, food service_models.Food) error {
	return nil
}

func (f *FoodRepo) UpdateFood(ctx context.Context, id string) error {
	return nil
}

func NewFoodRepo(db *mongo.Database) Food {
	return &FoodRepo{
		collection: db.Collection("food"),
	}
}
