package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/dto"
)

func (controller *Controller) jobList(c *fiber.Ctx) error {

	responses, err := controller.Service.JobList()
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

func (controller *Controller) jobGet(c *fiber.Ctx) error {

	jobId := c.Params("id")
	response, err := controller.Service.JobGet(jobId)
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

func (controller *Controller) jobCreate(c *fiber.Ctx) error {

	request := c.Body()
	response, err := controller.Service.JobCreate(request)
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
