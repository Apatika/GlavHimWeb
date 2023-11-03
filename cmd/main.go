package main

import (
	"apatikaWebServer/internal/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Println(err)
	}
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	srv := server.New()
	srv.Run()
}
