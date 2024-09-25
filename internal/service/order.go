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
	return o.orderRepo.GetOrder(ctx, id)
}

func (o *orderService) GetOrders(ctx context.Context) ([]service_models.Order, error) {
	return o.orderRepo.GetOrders(ctx)
}

func (o *orderService) CreateOrder(ctx context.Context, menu service_models.Order) error {
	return o.orderRepo.CreateOrder(ctx, menu)
}

func (o *orderService) UpdateOrder(ctx context.Context, id string) error {
	return o.orderRepo.UpdateOrder(ctx, id)
}

func NewOrderService(orderRepo repository.Order) Order {
	return &orderService{
		orderRepo: orderRepo,
	}
}
