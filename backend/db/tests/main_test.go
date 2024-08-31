package db_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-envconfig"
	"leogsouza.dev/superfin/config"
	db "leogsouza.dev/superfin/db/sqlc"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	ctx := context.Background()

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error on load env variables ", err)
	}

	var c config.TestConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Config.DBConfig.User, c.Config.DBConfig.Password, c.Config.DBConfig.Host, c.Config.DBConfig.Port, c.Config.DBConfig.DbName))

	if err != nil {
		log.Fatal("Could not connect to database", err)
	}

	testQuery = db.New(conn)

	os.Exit(m.Run())

}
