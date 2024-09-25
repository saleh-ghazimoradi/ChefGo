package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ChefGo/internal/service"
)

type invoiceHandler struct {
	invoiceService service.Invoice
}

func (i *invoiceHandler) GetInvoice(c *fiber.Ctx) error {
	return nil
}

func (i *invoiceHandler) GetInvoices(c *fiber.Ctx) error {
	return nil
}

func (i *invoiceHandler) CreateInvoice(c *fiber.Ctx) error {
	return nil
}

func (i *invoiceHandler) UpdateInvoice(c *fiber.Ctx) error {
	return nil
}

func NewInvoiceHandler(invoiceService service.Invoice) *invoiceHandler {
	return &invoiceHandler{
		invoiceService: invoiceService,
	}
}
