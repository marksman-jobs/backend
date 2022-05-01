package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/model"
	"github.com/marksman-jobs/backend/service"
)

type CandidateController struct {
	CandidateService service.CandidateService
}

func NewCandidateController(candidateService *service.CandidateService) CandidateController {
	return CandidateController{CandidateService: *candidateService}
}

func (controller *CandidateController) Route(app *fiber.App) {

	app.Get("/api/v1/candidates", controller.list)
	app.Get("/api/v1/candidates/:id", controller.get)
	app.Post("api/v1/candidates", controller.create)

}

func (controller *CandidateController) list(c *fiber.Ctx) error {

	responses, err := controller.CandidateService.List()
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

func (controller *CandidateController) get(c *fiber.Ctx) error {

	candidateId := c.Params("id")
	response, err := controller.CandidateService.Get(candidateId)
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

func (controller *CandidateController) create(c *fiber.Ctx) error {

	request := c.Body()
	response, err := controller.CandidateService.Create(request)
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
