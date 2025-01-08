package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ResponseParam struct {
	Data    any
	Message string
	Status  int
	Error   error
	PaginationMetaData
}

type PaginationMetaData struct {
	Total int
}

func SendResponse(c *fiber.Ctx, params ResponseParam) error {
	if params.Status == 0 {
		params.Status = 200
	}
	if params.Message == "" {
		params.Message = "Ok"
	}
	err := c.Status(params.Status).JSON(map[string]any{"data": params.Data, "success": true, "message": params.Message})
	if err != nil {
		SendErrorResponse(c, ResponseParam{Error: err})
	}
	return err
}

func SendErrorResponse(c *fiber.Ctx, params ResponseParam) error {
	log.Println("Error:", params.Error, params.Message)
	if params.Status == 0 {
		params.Status = 500
	}
	errorMessage := params.Error.Error()
	if errorMessage == "" {
		errorMessage = "Internal Server Error"
	}
	err := c.Status(params.Status).JSON(map[string]any{"success": false, "message": errorMessage})
	if err != nil {
		log.Println("Error :", err.Error())
	}
	return err
}
