package routes

import (
	"cfapp/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is working",
		})
	})

	apiEndpoint := app.Group("/api")
	apiVersionOne := apiEndpoint.Group("/v1")

	testEndpoint := apiVersionOne.Group("/test")
	testEndpoint.Get("/one", controller.CtrlOne)
}
