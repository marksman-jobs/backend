package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/config"
	"github.com/marksman-jobs/backend/service"
)

func InitializeServer(conf *config.Config, db *config.Databases) {

	service := service.NewService(*db.Postgres)
	controller := NewController(&service)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	controller.AttachRoutesToApp(app)

	fmt.Println("Server Initialized")

	app.Listen(":3000")

}
