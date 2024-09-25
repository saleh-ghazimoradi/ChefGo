package service

import (
	"context"
	"github.com/saleh-ghazimoradi/ChefGo/internal/repository"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
)

type Food interface {
	GetFood(ctx context.Context, id string) (service_models.Food, error)
	GetFoods(ctx context.Context) ([]service_models.Food, error)
	CreateFood(ctx context.Context, food service_models.Food) error
	UpdateFood(ctx context.Context, id string) error
}

type foodService struct {
	foodRepo repository.Food
}

func (f *foodService) GetFood(ctx context.Context, id string) (service_models.Food, error) {
	return service_models.Food{}, nil
}

func (f *foodService) GetFoods(ctx context.Context) ([]service_models.Food, error) {
	return nil, nil
}

func (f *foodService) CreateFood(ctx context.Context, food service_models.Food) error {
	return nil
}

func (f *foodService) UpdateFood(ctx context.Context, id string) error {
	return nil
}

func NewFoodService(foodRepo repository.Food) Food {
	return &foodService{
		foodRepo: foodRepo,
	}
}
