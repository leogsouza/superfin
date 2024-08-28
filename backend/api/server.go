package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewServer(port int) {
	f := fiber.New()

	f.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mesage": "Welcome to SuperFin",
		})
	})

	f.Listen(fmt.Sprintf(":%d", port))
}
