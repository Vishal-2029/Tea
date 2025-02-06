package main

import (
	"fmt"
	db "teamaking/config"
	"teamaking/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.Connect()
	fmt.Println("Let's Make a Tea..")

	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app) // Set up routes
	if err := app.Listen(":3030"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
