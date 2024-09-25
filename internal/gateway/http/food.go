package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ChefGo/internal/service"
)

type foodHandler struct {
	foodService service.Food
}

func (f *foodHandler) GetFood(c *fiber.Ctx) error {
	return nil
}

func (f *foodHandler) GetFoods(c *fiber.Ctx) error {
	return nil
}

func (f *foodHandler) CreateFood(c *fiber.Ctx) error {
	return nil
}

func (f *foodHandler) UpdateFood(c *fiber.Ctx) error {
	return nil
}

func NewFoodHandler(foodService service.Food) *foodHandler {
	return &foodHandler{
		foodService: foodService,
	}
}
