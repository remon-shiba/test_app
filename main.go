package main

import (
	"cfapp/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.AppRoutes(app)

	app.Listen(":4343")
}
