package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/config"
)

func InitializeServer(conf *config.Config) {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println("Server Initialized")

	app.Listen(":3000")

}
