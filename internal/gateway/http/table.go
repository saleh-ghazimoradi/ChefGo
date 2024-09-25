package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ChefGo/internal/service"
)

type tableHandler struct {
	tableService service.Table
}

func (t *tableHandler) GetTable(c *fiber.Ctx) error {
	return nil
}

func (t *tableHandler) GetTables(c *fiber.Ctx) error {
	return nil
}

func (t *tableHandler) CreateTable(c *fiber.Ctx) error {
	return nil
}

func (t *tableHandler) UpdateTable(c *fiber.Ctx) error {
	return nil
}

func NewTableHandler(tableService service.Table) *tableHandler {
	return &tableHandler{
		tableService: tableService,
	}
}
