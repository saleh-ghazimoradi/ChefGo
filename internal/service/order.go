package service

import (
	"context"
	"github.com/saleh-ghazimoradi/ChefGo/internal/repository"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
)

type Order interface {
	GetOrder(ctx context.Context, id string) (service_models.Order, error)
	GetOrders(ctx context.Context) ([]service_models.Order, error)
	CreateOrder(ctx context.Context, menu service_models.Order) error
	UpdateOrder(ctx context.Context, id string) error
}

type orderService struct {
	orderRepo repository.Order
}

func (o *orderService) GetOrder(ctx context.Context, id string) (service_models.Order, error) {
	return service_models.Order{}, nil
}

func (o *orderService) GetOrders(ctx context.Context) ([]service_models.Order, error) {
	return nil, nil
}

func (o *orderService) CreateOrder(ctx context.Context, menu service_models.Order) error {
	return nil
}

func (o *orderService) UpdateOrder(ctx context.Context, id string) error {
	return nil
}

func NewOrderService(orderRepo repository.Order) Order {
	return &orderService{
		orderRepo: orderRepo,
	}
}
