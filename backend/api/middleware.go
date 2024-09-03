package api

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"leogsouza.dev/superfin/utils"
)

func jwtAuthentication(c *fiber.Ctx) error {
	authToken := c.Get("Authorization")
	if authToken == "" {
		return fmt.Errorf("unauthorized")
	}
	tokenSplit := strings.Split(authToken, " ")
	if len(tokenSplit) < 2 {
		return fmt.Errorf("unauthorized")
	}
	token := tokenSplit[1]

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
