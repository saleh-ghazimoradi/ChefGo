package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Invoice interface {
	GetInvoice(ctx context.Context, id string) (service_models.Invoice, error)
	GetInvoices(ctx context.Context) ([]service_models.Invoice, error)
	CreateInvoice(ctx context.Context, invoice service_models.Invoice) error
	UpdateInvoice(ctx context.Context, id string) error
}

type InvoiceRepo struct {
	collection *mongo.Collection
}

func (i *InvoiceRepo) GetInvoice(ctx context.Context, id string) (service_models.Invoice, error) {
	return service_models.Invoice{}, nil
}

func (i *InvoiceRepo) GetInvoices(ctx context.Context) ([]service_models.Invoice, error) {
	return nil, nil
}

func (i *InvoiceRepo) CreateInvoice(ctx context.Context, invoice service_models.Invoice) error {
	return nil
}

func (i *InvoiceRepo) UpdateInvoice(ctx context.Context, id string) error {
	return nil
}

func NewInvoiceRepo(db *mongo.Database) Invoice {
	return &InvoiceRepo{
		collection: db.Collection("invoice"),
	}
}
