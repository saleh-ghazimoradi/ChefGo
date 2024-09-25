package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order interface {
	GetOrder(ctx context.Context, id string) (service_models.Order, error)
	GetOrders(ctx context.Context) ([]service_models.Order, error)
	CreateOrder(ctx context.Context, menu service_models.Order) error
	UpdateOrder(ctx context.Context, id string) error
}

type orderRepo struct {
	collection *mongo.Collection
}

func (o *orderRepo) GetOrder(ctx context.Context, id string) (service_models.Order, error) {
	return service_models.Order{}, nil
}

func (o *orderRepo) GetOrders(ctx context.Context) ([]service_models.Order, error) {
	return nil, nil
}

func (o *orderRepo) CreateOrder(ctx context.Context, menu service_models.Order) error {
	return nil
}

func (o *orderRepo) UpdateOrder(ctx context.Context, id string) error {
	return nil
}

func NewOrderRepo(db *mongo.Database) Order {
	return &orderRepo{
		collection: db.Collection("order"),
	}
}
