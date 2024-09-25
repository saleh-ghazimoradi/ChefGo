package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ChefGo/internal/service"
)

type orderItems struct {
	orderItemsService service.OrderItem
}

func (o *orderItems) GetOrderItem(c *fiber.Ctx) error {
	return nil
}

func (o *orderItems) GetOrderItems(c *fiber.Ctx) error {
	return nil
}

func (o *orderItems) GetOrderItemsByOrder(c *fiber.Ctx) error {
	return nil
}

func (o *orderItems) ItemsByOrder(c *fiber.Ctx) error {
	return nil
}

func (o *orderItems) CreateOrderItem(c *fiber.Ctx) error {
	return nil
}

func (o *orderItems) UpdateOrderItem(c *fiber.Ctx) error {
	return nil
}

func NewOrderItems(orderItemsService service.OrderItem) *orderItems {
	return &orderItems{
		orderItemsService: orderItemsService,
	}
}
