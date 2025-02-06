package routes

import (
	"teamaking/controllers"

	"github.com/gofiber/fiber/v2" // Correct import path
)

func Setup(app *fiber.App) {
	app.Get("/Tea", controllers.ReadTea) // Changed from "/Tea/" to "/Tea"
	app.Post("/Tea", controllers.CreateTea)
	app.Delete("/Tea/:TeaID", controllers.DeleteTea)
	app.Put("/Tea/:TeaID", controllers.UpdateTea)
}
