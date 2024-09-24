package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order interface {
	GetMenu(ctx context.Context, id string) (service_models.Menu, error)
	GetMenus(ctx context.Context) ([]service_models.Menu, error)
	CreateMenu(ctx context.Context, menu service_models.Menu) error
	UpdateMenu(ctx context.Context, id string) error
}

type orderRepo struct {
	collection *mongo.Collection
}

func (o *orderRepo) GetMenu(ctx context.Context, id string) (service_models.Menu, error) {
	return service_models.Menu{}, nil
}

func (o *orderRepo) GetMenus(ctx context.Context) ([]service_models.Menu, error) {
	return nil, nil
}

func (o *orderRepo) CreateMenu(ctx context.Context, menu service_models.Menu) error {
	return nil
}

func (o *orderRepo) UpdateMenu(ctx context.Context, id string) error {
	return nil
}

func NewOrderRepo(db *mongo.Database) Order {
	return &orderRepo{
		collection: db.Collection("order"),
	}
}
