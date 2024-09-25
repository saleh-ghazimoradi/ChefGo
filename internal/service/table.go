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
	return t.tableRepo.GetTable(ctx, id)
}

func (t *tableService) GetTables(ctx context.Context) ([]service_models.Table, error) {
	return t.tableRepo.GetTables(ctx)
}

func (t *tableService) CreateTable(ctx context.Context, menu service_models.Table) error {
	return t.tableRepo.CreateTable(ctx, menu)
}

func (t *tableService) UpdateTable(ctx context.Context, id string) error {
	return t.tableRepo.UpdateTable(ctx, id)
}

func NewTableService(tableRepo repository.Table) Table {
	return &tableService{
		tableRepo: tableRepo,
	}
}
