package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nandono206/FinalSeasonPart3/backend/database"
	"github.com/nandono206/FinalSeasonPart3/backend/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

func main() {

	database.Connect()
	


    app := fiber.New()
	app.Use(cors.New())

    routes.Initialize(app)

	

    app.Listen(":8000")
}