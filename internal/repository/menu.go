package repository

import (
	"context"
	service_models "github.com/saleh-ghazimoradi/ChefGo/internal/service/service-models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Menu interface {
	GetMenu(ctx context.Context, id string) (service_models.Menu, error)
	GetMenus(ctx context.Context) ([]service_models.Menu, error)
	CreateMenu(ctx context.Context, menu service_models.Menu) error
	UpdateMenu(ctx context.Context, id string) error
}

type menuRepo struct {
	collection *mongo.Collection
}

func (m *menuRepo) GetMenu(ctx context.Context, id string) (service_models.Menu, error) {
	return service_models.Menu{}, nil
}

func (m *menuRepo) GetMenus(ctx context.Context) ([]service_models.Menu, error) {
	return nil, nil
}

func (m *menuRepo) CreateMenu(ctx context.Context, menu service_models.Menu) error {
	return nil
}

func (m *menuRepo) UpdateMenu(ctx context.Context, id string) error {
	return nil
}

func NewMenu(db *mongo.Database) Menu {
	return &menuRepo{
		collection: db.Collection("menu"),
	}
}
