package controller

import (
	responseUtils "github.com/chiragthapa777/wishlist-api/utils/response"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return responseUtils.SendResponse(c, responseUtils.ResponseParam{
		Message: "Hello from server",
	})
}
