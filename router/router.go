package router

import (
	authController "github.com/chiragthapa777/wishlist-api/controller/auth"
	healthControllers "github.com/chiragthapa777/wishlist-api/controller/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Logger setup
	api := app.Group("/api", logger.New())

	// Health
	api.Get("/", healthControllers.Hello)

	// auth
	authGroup := api.Group("/auth", logger.New())
	authGroup.Post("/register", authController.Register)
	authGroup.Post("/login", authController.Login)
	authGroup.Get("/", authController.GetCurrentUser)
}
