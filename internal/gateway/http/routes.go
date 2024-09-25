package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ChefGo/config"
	"github.com/saleh-ghazimoradi/ChefGo/internal/repository"
	"github.com/saleh-ghazimoradi/ChefGo/internal/service"
	"github.com/saleh-ghazimoradi/ChefGo/utils"
)

func registerRoutes(app *fiber.App) {
	dbClient, err := utils.ConnectMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	db := dbClient.Database(config.AppConfig.Databases.MongoDB.DatabaseName)

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := NewUserHandler(userService)

	foodRepo := repository.NewFoodRepo(db)
	foodService := service.NewFoodService(foodRepo)
	foodController := NewFoodHandler(foodService)

	tableRepo := repository.NewTableRepo(db)
	tableService := service.NewTableService(tableRepo)
	tableController := NewTableHandler(tableService)

	invoiceRepo := repository.NewInvoiceRepo(db)
	invoiceService := service.NewInvoiceService(invoiceRepo)
	invoiceController := NewInvoiceHandler(invoiceService)

	menuRepo := repository.NewMenu(db)
	menuService := service.NewMenuService(menuRepo)
	menuController := NewMenuHandler(menuService)

	orderRepo := repository.NewOrderRepo(db)
	orderService := service.NewOrderService(orderRepo)
	orderController := NewOrderHandler(orderService)

	orderItemRepo := repository.NewOrderItemRepo(db)
	orderItemService := service.NewOrderItemService(orderItemRepo)
	orderItemController := NewOrderItemsHandler(orderItemService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Chef Go Gateway")
	})

	v1 := app.Group("/v1")

	/*---------------------- User Routes ---------------------*/
	v1.Get("/users", userController.GetUsers)
	v1.Get("/users/:user_id", userController.GetUser)
	v1.Post("/users/signup", userController.SignUp)
	v1.Post("/users/login", userController.Login)

	/*---------------------- Food Routes ---------------------*/
	v1.Get("/foods", foodController.GetFoods)
	v1.Get("/foods/:food_id", foodController.GetFood)
	v1.Post("/foods", foodController.CreateFood)
	v1.Patch("/foods/:food_id", foodController.UpdateFood)

	/*---------------------- Invoice Routes ---------------------*/
	v1.Get("/invoices", invoiceController.GetInvoices)
	v1.Get("/invoices/:invoice_id", invoiceController.GetInvoice)
	v1.Post("/invoices", invoiceController.CreateInvoice)
	v1.Patch("/invoices/:invoice_id", invoiceController.UpdateInvoice)

	/*---------------------- Menu Routes ---------------------*/
	v1.Get("/menus", menuController.GetMenus)
	v1.Get("/menus/:menu_id", menuController.GetMenu)
	v1.Post("/menus", menuController.CreateMenu)
	v1.Patch("/menus/:menu_id", menuController.UpdateMenu)

	/*---------------------- Order Routes ---------------------*/
	v1.Get("/orders", orderController.GetOrders)
	v1.Get("/orders/:order_id", orderController.GetOrder)
	v1.Post("/orders", orderController.CreateOrder)
	v1.Patch("/orders/:order_id", orderController.UpdateOrder)

	/*---------------------- Order Items Routes ---------------------*/
	v1.Get("/orderItems", orderItemController.GetOrderItems)
	v1.Get("/orderItems/:orderItem_id", orderItemController.GetOrderItem)
	v1.Get("/orderItems-order/:order_id", orderItemController.GetOrderItemsByOrder)
	v1.Post("/orderItems", orderItemController.CreateOrderItem)
	v1.Patch("/orderItems/:orderItem_id", orderItemController.UpdateOrderItem)

	/*---------------------- Table Routes ---------------------*/
	v1.Get("/tables", tableController.GetTables)
	v1.Get("/tables/:table_id", tableController.GetTables)
	v1.Post("/tables", tableController.CreateTable)
	v1.Patch("/tables/:table_id", tableController.UpdateTable)
}
