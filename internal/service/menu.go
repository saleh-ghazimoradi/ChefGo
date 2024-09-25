package service

import (
	"context"
	"github.com/saleh-ghazimoradi/ChefGo/internal/repository"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
)

type Menu interface {
	GetMenu(ctx context.Context, id string) (service_models.Menu, error)
	GetMenus(ctx context.Context) ([]service_models.Menu, error)
	CreateMenu(ctx context.Context, menu service_models.Menu) error
	UpdateMenu(ctx context.Context, id string) error
}

type menuService struct {
	menuRepo repository.Menu
}

func (m *menuService) GetMenu(ctx context.Context, id string) (service_models.Menu, error) {
	return service_models.Menu{}, nil
}

func (m *menuService) GetMenus(ctx context.Context) ([]service_models.Menu, error) {
	return nil, nil
}

func (m *menuService) CreateMenu(ctx context.Context, menu service_models.Menu) error {
	return nil
}

func (m *menuService) UpdateMenu(ctx context.Context, id string) error {
	return nil
}

func NewMenuService(menuRepo repository.Menu) Menu {
	return &menuService{
		menuRepo: menuRepo,
	}
}
