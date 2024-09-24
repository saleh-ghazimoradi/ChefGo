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

type invoiceRepo struct {
	collection *mongo.Collection
}

func (i *invoiceRepo) GetInvoice(ctx context.Context, id string) (service_models.Invoice, error) {
	return service_models.Invoice{}, nil
}

func (i *invoiceRepo) GetInvoices(ctx context.Context) ([]service_models.Invoice, error) {
	return nil, nil
}

func (i *invoiceRepo) CreateInvoice(ctx context.Context, invoice service_models.Invoice) error {
	return nil
}

func (i *invoiceRepo) UpdateInvoice(ctx context.Context, id string) error {
	return nil
}

func NewInvoiceRepo(db *mongo.Database) Invoice {
	return &invoiceRepo{
		collection: db.Collection("invoice"),
	}
}
