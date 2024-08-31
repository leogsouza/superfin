package api

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	db "leogsouza.dev/superfin/db/sqlc"
	"leogsouza.dev/superfin/utils"
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
	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  10,
	}

	users, err := u.server.queries.ListUsers(c.Context(), arg)
	if err != nil {
		return err
	}

	usersResponse := []*userResponse{}

	for _, user := range users {
		ur := transformDbUsertoUserResponse(&user)
		usersResponse = append(usersResponse, ur)
	}

	return c.JSON(usersResponse)
}

type createUserParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type userResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *userHandler) createUser(c *fiber.Ctx) error {
	var userParams createUserParams

	// parse body content into struct
	if err := c.BodyParser(&userParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// validate the data
	validate := validator.New()
	if err := validate.Struct(userParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	hashedPassword, err := utils.GenerateHashPassword(userParams.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userDb := db.CreateUserParams{}
	userDb.Email = userParams.Email
	userDb.Password = hashedPassword
	user, err := u.server.queries.CreateUser(c.Context(), userDb)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	userResponse := transformDbUsertoUserResponse(&user)

	return c.Status(fiber.StatusOK).JSON(userResponse)
}

func transformDbUsertoUserResponse(dbUser *db.User) *userResponse {
	return &userResponse{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}
