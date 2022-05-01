package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/model"
	"github.com/marksman-jobs/backend/service"
)

type CompanyController struct {
	CompanyService service.CompanyService
}

func NewCompanyController(companyService *service.CompanyService) CompanyController {
	return CompanyController{CompanyService: *companyService}
}

func (controller *CompanyController) Route(app *fiber.App) {

	app.Get("/api/v1/company", controller.list)
	app.Get("/api/v1/company/:id", controller.get)
	app.Post("api/v1/company", controller.create)

}

func (controller *CompanyController) list(c *fiber.Ctx) error {

	responses, err := controller.CompanyService.List()
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})

}

func (controller *CompanyController) get(c *fiber.Ctx) error {

	companyId := c.Params("id")
	response, err := controller.CompanyService.Get(companyId)
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

func (controller *CompanyController) create(c *fiber.Ctx) error {

	request := c.Body()
	response, err := controller.CompanyService.Create(request)
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
