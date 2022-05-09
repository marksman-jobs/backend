package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marksman-jobs/backend/dto"
)

func (controller *Controller) companyList(c *fiber.Ctx) error {

	responses, err := controller.Service.CompanyList()
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

func (controller *Controller) companyGet(c *fiber.Ctx) error {

	companyId := c.Params("id")
	response, err := controller.Service.CompanyGet(companyId)
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

func (controller *Controller) companyCreate(c *fiber.Ctx) error {

	request := c.Body()
	response, err := controller.Service.CompanyCreate(request)
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
