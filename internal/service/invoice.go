package service

import (
	"context"
	"github.com/saleh-ghazimoradi/ChefGo/internal/repository"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
)

type Invoice interface {
	GetInvoice(ctx context.Context, id string) (service_models.Invoice, error)
	GetInvoices(ctx context.Context) ([]service_models.Invoice, error)
	CreateInvoice(ctx context.Context, invoice service_models.Invoice) error
	UpdateInvoice(ctx context.Context, id string) error
}

type invoiceService struct {
	invoiceRepo repository.Invoice
}

func (i *invoiceService) GetInvoice(ctx context.Context, id string) (service_models.Invoice, error) {
	return service_models.Invoice{}, nil
}

func (i *invoiceService) GetInvoices(ctx context.Context) ([]service_models.Invoice, error) {
	return nil, nil
}

func (i *invoiceService) CreateInvoice(ctx context.Context, invoice service_models.Invoice) error {
	return nil
}

func (i *invoiceService) UpdateInvoice(ctx context.Context, id string) error {
	return nil
}

func NewInvoiceService(invoiceRepo repository.Invoice) Invoice {
	return &invoiceService{
		invoiceRepo: invoiceRepo,
	}
}
