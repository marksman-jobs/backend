package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/config"
	"github.com/marksman-jobs/backend/repository"
	"github.com/marksman-jobs/backend/service"
)

func InitializeServer(conf *config.Config, db *config.Databases) {

	candidateRepository := repository.NewCandidateRepository(db.Postgres)
	candidateService := service.NewCandidateService(&candidateRepository)
	candidateController := NewCandidateController(&candidateService)

	companyRepository := repository.NewCompanyRepository(db.Postgres)
	companyService := service.NewCompanyService(&companyRepository)
	companyController := NewCompanyController(&companyService)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	candidateController.Route(app)
	companyController.Route(app)

	fmt.Println("Server Initialized")

	app.Listen(":3000")

}
