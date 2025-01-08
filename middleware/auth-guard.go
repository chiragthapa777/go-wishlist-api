package middleware

import (
	"errors"
	"log"
	"strings"

	authService "github.com/chiragthapa777/wishlist-api/service/auth"
	responseUtils "github.com/chiragthapa777/wishlist-api/utils/response"
	"github.com/gofiber/fiber/v2"
)

func AuthGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: errors.New("not authorized"), Status: 401})
		}

		stringArray := strings.Split(authorization, " ")

		if len(stringArray) < 2 {
			return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: errors.New("not authorized"), Status: 401})
		}

		jwtToken := stringArray[1]

		payload, err := authService.ValidateJWT(jwtToken)
		if err != nil || payload == nil {
			log.Println("err", err.Error())
			return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: errors.New("not authorized"), Status: 401})
		}

		userID := payload.UserId

		c.Locals("userId", userID)
		return c.Next()
	}
}
