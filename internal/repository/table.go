package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Table interface {
	GetTable(ctx context.Context, id string) (service_models.Table, error)
	GetTables(ctx context.Context) ([]service_models.Table, error)
	CreateTable(ctx context.Context, menu service_models.Table) error
	UpdateTable(ctx context.Context, id string) error
}

type tableRepo struct {
	collection *mongo.Collection
}

func (t *tableRepo) GetTable(ctx context.Context, id string) (service_models.Table, error) {
	return service_models.Table{}, nil
}

func (t *tableRepo) GetTables(ctx context.Context) ([]service_models.Table, error) {
	return nil, nil
}

func (t *tableRepo) CreateTable(ctx context.Context, menu service_models.Table) error {
	return nil
}

func (t *tableRepo) UpdateTable(ctx context.Context, id string) error {
	return nil
}

func NewTableRepo(db *mongo.Database) Table {
	return &tableRepo{
		collection: db.Collection("table"),
	}
}
