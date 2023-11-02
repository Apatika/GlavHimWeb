package main

import (
	"apatikaWebServer/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	cfg := config.New()
	log.Println(cfg.Server.Port)
}
