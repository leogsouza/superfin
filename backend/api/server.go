package api

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"leogsouza.dev/superfin/config"
	db "leogsouza.dev/superfin/db/sqlc"
)

type Server struct {
	queries *db.Queries
	router  *fiber.App
}

func NewServer(c *config.Config) *Server {

	conn, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.DBConfig.User, c.DBConfig.Password, c.DBConfig.Host, c.DBConfig.Port, c.DBConfig.DbName))
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %+v", err))
	}

	q := db.New(conn)
	f := fiber.New()

	return &Server{
		queries: q,
		router:  f,
	}
}

func (s *Server) Start(port int) {

	s.router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mesage": "Welcome to SuperFin",
		})
	})

	s.router.Listen(fmt.Sprintf(":%d", port))
}
