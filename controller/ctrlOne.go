package controller

import "github.com/gofiber/fiber/v2"

func CtrlOne(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Message": "This is controller",
	})
}
