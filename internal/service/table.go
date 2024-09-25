package service

import (
	"context"
	"github.com/saleh-ghazimoradi/ChefGo/internal/repository"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
)

type Table interface {
	GetTable(ctx context.Context, id string) (service_models.Table, error)
	GetTables(ctx context.Context) ([]service_models.Table, error)
	CreateTable(ctx context.Context, menu service_models.Table) error
	UpdateTable(ctx context.Context, id string) error
}

type tableService struct {
	tableRepo repository.Table
}

func (t *tableService) GetTable(ctx context.Context, id string) (service_models.Table, error) {
	return service_models.Table{}, nil
}

func (t *tableService) GetTables(ctx context.Context) ([]service_models.Table, error) {
	return nil, nil
}

func (t *tableService) CreateTable(ctx context.Context, menu service_models.Table) error {
	return nil
}

func (t *tableService) UpdateTable(ctx context.Context, id string) error {
	return nil
}

func NewTableService(tableRepo repository.Table) Table {
	return &tableService{
		tableRepo: tableRepo,
	}
}
