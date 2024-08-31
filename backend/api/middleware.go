package api

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"leogsouza.dev/superfin/utils"
)

func jwtAuthentication(c *fiber.Ctx) error {
	token := c.Get("X-Api-Token")
	if token == "" {
		return fmt.Errorf("unauthorized")
	}

	jwtWrapper := utils.JwtWrapper{
		SecretKey: os.Getenv("APP_SECRET_KEY"),
		Issuer:    "AuthUser",
	}

	claim, err := jwtWrapper.ValidateToken(token)

	if err != nil {
		return err
	}

	c.Locals("userEmail", claim.Email)

	return c.Next()
}
