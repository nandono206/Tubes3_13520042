package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/nandono206/FinalSeasonPart3/backend/controller"
	"github.com/nandono206/FinalSeasonPart3/backend/database"
	"github.com/nandono206/FinalSeasonPart3/backend/routes"
)

func main() {

	database.Connect()
	


    app := fiber.New()
	app.Use(cors.New())

    routes.Initialize(app)
	// controller.SetupRoutes()

	

    app.Listen(":8000")
}