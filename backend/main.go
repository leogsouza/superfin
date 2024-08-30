package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"leogsouza.dev/superfin/api"
	"leogsouza.dev/superfin/config"
)

func main() {
	godotenv.Load()
	var c config.Config
	if err := envconfig.Process(context.Background(), &c); err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(&c)

	server.Start(c.AppConfig.Port)
}
