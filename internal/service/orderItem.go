package service

import (
	"context"
	"github.com/saleh-ghazimoradi/ChefGo/internal/repository"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
)

type OrderItem interface {
	GetOrderItem(ctx context.Context, id string) (service_models.OrderItem, error)
	GetOrderItems(ctx context.Context) ([]service_models.OrderItem, error)
	CreateOrderItem(ctx context.Context, menu service_models.OrderItem) error
	UpdateOrderItem(ctx context.Context, id string) error
}

type orderItemService struct {
	orderItemRepo repository.OrderItem
}

func (o *orderItemService) GetOrderItem(ctx context.Context, id string) (service_models.OrderItem, error) {
	return o.orderItemRepo.GetOrderItem(ctx, id)
}

func (o *orderItemService) GetOrderItems(ctx context.Context) ([]service_models.OrderItem, error) {
	return o.orderItemRepo.GetOrderItems(ctx)
}

func (o *orderItemService) CreateOrderItem(ctx context.Context, menu service_models.OrderItem) error {
	return o.orderItemRepo.CreateOrderItem(ctx, menu)
}

func (o *orderItemService) UpdateOrderItem(ctx context.Context, id string) error {
	return o.orderItemRepo.UpdateOrderItem(ctx, id)
}

func NewOrderItemService(orderItemRepo repository.OrderItem) OrderItem {
	return &orderItemService{
		orderItemRepo: orderItemRepo,
	}
}
