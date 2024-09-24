package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItem interface {
	GetOrderItem(ctx context.Context, id string) (service_models.OrderItem, error)
	GetOrderItems(ctx context.Context) ([]service_models.OrderItem, error)
	CreateOrderItem(ctx context.Context, menu service_models.OrderItem) error
	UpdateOrderItem(ctx context.Context, id string) error
}

type orderItemRepo struct {
	collection *mongo.Collection
}

func (o *orderItemRepo) GetOrderItem(ctx context.Context, id string) (service_models.OrderItem, error) {
	return service_models.OrderItem{}, nil
}

func (o *orderItemRepo) GetOrderItems(ctx context.Context) ([]service_models.OrderItem, error) {
	return nil, nil
}

func (o *orderItemRepo) CreateOrderItem(ctx context.Context, menu service_models.OrderItem) error {
	return nil
}

func (o *orderItemRepo) UpdateOrderItem(ctx context.Context, id string) error {
	return nil
}

func NewOrderItemRepo(db *mongo.Database) OrderItem {
	return &orderItemRepo{
		collection: db.Collection("order_item"),
	}
}
