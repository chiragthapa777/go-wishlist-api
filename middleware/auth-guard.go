package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		// Extract user ID from the token (mocked here; implement your own logic)
		userID := "12345" // Replace this with actual extraction logic

		// Store the user ID in the context
		c.Locals("userID", userID)
		log.Println(authHeader)
		return c.Next()
	}
}
