package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/service"
)

type Controller struct {
	Service service.Service
}

func NewController(service *service.Service) Controller {
	return Controller{
		Service: *service,
	}
}

func (controller *Controller) AttachRoutesToApp(app *fiber.App) {

	app.Get("/api/v1/candidates", controller.candidateList)
	app.Get("/api/v1/candidates/:id", controller.candidateGet)
	app.Post("api/v1/candidates", controller.candidateCreate)

	app.Get("/api/v1/companies", controller.companyList)
	app.Get("/api/v1/companies/:id", controller.companyGet)
	app.Post("api/v1/companies", controller.companyCreate)

	app.Get("/api/v1/jobs", controller.jobList)
	app.Get("/api/v1/jobs/:id", controller.jobGet)
	app.Post("api/v1/jobs", controller.jobCreate)

}
