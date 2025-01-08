package main

import (
	"log"

	"github.com/chiragthapa777/wishlist-api/config"
	"github.com/chiragthapa777/wishlist-api/database"
	"github.com/chiragthapa777/wishlist-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// load config
	config.LoadConfig()

	// connect to db
	database.ConnectDB()

	// create app
	app := fiber.New()
	app.Use(cors.New())

	// setup routes
	router.SetupRoutes(app)

	log.Println("Test long", config.GetConfig())

	serverBaseUrl := config.GetConfig().Host + ":" + config.GetConfig().Port

	log.Printf("Server is running at %s \n", serverBaseUrl)

	log.Fatal(app.Listen(serverBaseUrl))
}
