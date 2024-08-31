package api

import (
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"leogsouza.dev/superfin/utils"
)

type authHandler struct {
	server *Server
}

func NewAuthHandler(server *Server) *authHandler {
	return &authHandler{
		server: server,
	}
}

func invalidCredentials(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(genericResponse{
		Type:    "error",
		Message: "invalid credentials",
	})
}

func unauthorizedUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(genericResponse{
		Type:    "error",
		Message: "unauthorized user",
	})
}

func (a *authHandler) RegisterRoutes() {
	authGroup := a.server.Router.Group("/auth")
	authGroup.Post("login", a.login)
}

func (a *authHandler) login(c *fiber.Ctx) error {
	validate := validator.New()
	var authParams authParams
	if err := c.BodyParser(&authParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(genericResponse{
			Type:    "error",
			Message: err.Error(),
		})
	}

	if err := validate.Struct(authParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(genericResponse{
			Type:    "error",
			Message: err.Error(),
		})
	}

	user, err := a.server.queries.GetUserByEmail(c.Context(), authParams.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return invalidCredentials(c)
		}
	}

	if !utils.VerifyPassword(user.Password, authParams.Password) {
		return invalidCredentials(c)
	}

	jwtwrapper := utils.JwtWrapper{
		SecretKey:       a.server.Config.AppConfig.SecretKey,
		Issuer:          "AuthUser",
		ExpirationHours: 4,
	}

	tokenStr, err := jwtwrapper.GenerateToken(user.Email)
	if err != nil {
		return unauthorizedUser(c)
	}

	userResponse := transformDbUsertoUserResponse(&user)

	return c.Status(fiber.StatusOK).JSON(&authResponse{
		User:  userResponse,
		Token: tokenStr,
	})
}
