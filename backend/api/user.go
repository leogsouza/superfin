package api

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	db "leogsouza.dev/superfin/db/sqlc"
	"leogsouza.dev/superfin/utils"
)

var (
	errIncorrectCredentials = errors.New("incorrect credentials")
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
	userGroup.Get("", jwtAuthentication, u.listUsers)
	userGroup.Get("me", jwtAuthentication, u.getLoggedUser)
	userGroup.Post("", u.createUser)
	userGroup.Put("/:id", jwtAuthentication, u.updateUser)

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

func (u *userHandler) getLoggedUser(c *fiber.Ctx) error {
	userEmail, ok := c.Locals("userEmail").(string)
	if !ok || userEmail == "" {
		return unauthorizedUser(c)
	}

	user, err := u.server.queries.GetUserByEmail(c.Context(), userEmail)
	if err != nil {
		return unauthorizedUser(c)
	}

	userResponse := transformDbUsertoUserResponse(&user)
	return c.Status(200).JSON(userResponse)
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

	return c.Status(fiber.StatusCreated).JSON(userResponse)
}

func (u userHandler) updateUser(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = u.server.queries.GetUserById(c.Context(), int64(userID))
	if err == sql.ErrNoRows {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errIncorrectCredentials.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var userParams updateUserParams
	if err := c.BodyParser(&userParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userParams.ID = int64(userID)
	userParams.UpdatedAt = time.Now()

	// validate the data
	validate := validator.New()
	if err := validate.Struct(userParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	hashedPassword, err := utils.GenerateHashPassword(userParams.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userDb := db.UpdateUserPasswordParams{
		ID:        userParams.ID,
		Password:  hashedPassword,
		UpdatedAt: userParams.UpdatedAt,
	}

	userUpdated, err := u.server.queries.UpdateUserPassword(c.Context(), userDb)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := transformDbUsertoUserResponse(&userUpdated)

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
