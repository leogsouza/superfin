package api

import (
	"github.com/gofiber/fiber/v2"
	db "leogsouza.dev/superfin/db/sqlc"
)

type userHandler struct {
	server *Server
}

func NewUserHandler(server *Server) *userHandler {
	return &userHandler{
		server: server,
	}
}

func (u *userHandler) RegisterRoutes() {
	userGroup := u.server.Router.Group("/users")
	userGroup.Get("", u.listUsers)
	userGroup.Post("", u.createUser)
}

func (u *userHandler) listUsers(c *fiber.Ctx) error {
	return c.JSON([]string{"user1", "user2"})
}

func (u *userHandler) createUser(c *fiber.Ctx) error {
	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  10,
	}

	users, err := u.server.queries.ListUsers(c.Context(), arg)
	if err != nil {
		return err
	}
	return c.JSON(users)
}
