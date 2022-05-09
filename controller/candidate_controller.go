package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/dto"
)

func (controller *Controller) candidateList(c *fiber.Ctx) error {

	responses, err := controller.Service.CandidateList()
	if err != nil {
		return c.JSON(dto.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.JSON(dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})

}

func (controller *Controller) candidateGet(c *fiber.Ctx) error {

	candidateId := c.Params("id")
	response, err := controller.Service.CandidateGet(candidateId)
	if err != nil {
		return c.JSON(dto.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.JSON(dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

func (controller *Controller) candidateCreate(c *fiber.Ctx) error {

	request := c.Body()
	response, err := controller.Service.CandidateCreate(request)
	if err != nil {
		return c.JSON(dto.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.JSON(dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
