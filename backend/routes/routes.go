package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nandono206/FinalSeasonPart3/backend/controller"

)

func Initialize(app *fiber.App){

	app.Post("/api/addPenyakit", controller.AddPenyakit)
	app.Post("/api/tesSakit", controller.IsSakit)
	app.Post("/api/history", controller.CekHistory)
	

}