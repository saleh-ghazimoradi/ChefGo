package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ChefGo/internal/service"
)

type orderHandler struct {
	orderService service.Order
}

func (o *orderHandler) GetOrder(c *fiber.Ctx) error {
	return nil
}

func (o *orderHandler) GetOrders(c *fiber.Ctx) error {
	return nil
}

func (o *orderHandler) CreateOrder(c *fiber.Ctx) error {
	return nil
}

func (o *orderHandler) UpdateOrder(c *fiber.Ctx) error {
	return nil
}

func NewOrderHandler(orderService service.Order) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}
