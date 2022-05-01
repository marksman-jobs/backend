package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/config"
	"github.com/marksman-jobs/backend/repository"
	"github.com/marksman-jobs/backend/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeServer(conf *config.Config, db *mongo.Database) {

	candidateRepository := repository.NewCandidateRepository(db)
	candidateService := service.NewCandidateService(&candidateRepository)
	candidateController := NewCandidateController(&candidateService)

	companyRepository := repository.NewCompanyRepository(db)
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
