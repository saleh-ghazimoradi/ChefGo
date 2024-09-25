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
	return m.menuRepo.GetMenu(ctx, id)
}

func (m *menuService) GetMenus(ctx context.Context) ([]service_models.Menu, error) {
	return m.menuRepo.GetMenus(ctx)
}

func (m *menuService) CreateMenu(ctx context.Context, menu service_models.Menu) error {
	return m.menuRepo.CreateMenu(ctx, menu)
}

func (m *menuService) UpdateMenu(ctx context.Context, id string) error {
	return m.menuRepo.UpdateMenu(ctx, id)
}

func NewMenuService(menuRepo repository.Menu) Menu {
	return &menuService{
		menuRepo: menuRepo,
	}
}
