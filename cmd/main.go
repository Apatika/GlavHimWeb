package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println(".env file not found, use default params")
	}
	//cfg := config.New()
}
