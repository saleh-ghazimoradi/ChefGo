package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ChefGo/internal/service"
)

type menuHandler struct {
	menuService service.Menu
}

func (m *menuHandler) GetMenu(c *fiber.Ctx) error {
	return nil
}

func (m *menuHandler) GetMenus(c *fiber.Ctx) error {
	return nil
}

func (m *menuHandler) CreateMenu(c *fiber.Ctx) error {
	return nil
}

func (m *menuHandler) UpdateMenu(c *fiber.Ctx) error {
	return nil
}

func NewMenuHandler(menuService service.Menu) *menuHandler {
	return &menuHandler{menuService: menuService}
}
