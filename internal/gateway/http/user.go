package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ChefGo/internal/service"
)

type userHandler struct {
	userService service.User
}

func (u *userHandler) GetUser(c *fiber.Ctx) error {
	return nil
}

func (u *userHandler) GetUsers(c *fiber.Ctx) error {
	return nil
}

func (u *userHandler) SignUp(c *fiber.Ctx) error {
	return nil
}

func (u *userHandler) Login(c *fiber.Ctx) error {
	return nil
}

func NewUserHandler(userService service.User) *userHandler {
	return &userHandler{
		userService: userService,
	}
}
