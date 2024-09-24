package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Food interface {
	GetFood(ctx context.Context, id string) (Food, error)
	GetAllFood(ctx context.Context) ([]Food, error)
	CreateFood(ctx context.Context, food service_models.Food) (service_models.Food, error)
	UpdateFood(ctx context.Context, id string, food service_models.Food) (service_models.Food, error)
}

type FoodRepo struct {
	collection *mongo.Collection
}

func (f *FoodRepo) GetAllFood(ctx context.Context) ([]Food, error) {
	return nil, nil
}

func (f *FoodRepo) GetFood(ctx context.Context, id string) (Food, error) {
	return nil, nil
}

func (f *FoodRepo) CreateFood(ctx context.Context, food service_models.Food) (service_models.Food, error) {
	return service_models.Food{}, nil
}

func (f *FoodRepo) UpdateFood(ctx context.Context, id string, food service_models.Food) (service_models.Food, error) {
	return service_models.Food{}, nil
}

func NewFoodRepo(db *mongo.Database) Food {
	return &FoodRepo{
		collection: db.Collection("food"),
	}
}
