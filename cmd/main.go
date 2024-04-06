package main

import (
	"context"
	"glavhim-app/internal/config"
	"glavhim-app/internal/server"
	"glavhim-app/internal/storage"
	handler "glavhim-app/internal/transport"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

var DB_URI string

func main() {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println("load environment variables")
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatalln("ERROR:", err)
	}

	log.Println("load config")
	err = config.New()
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}

	log.Println("get http handler")
	handler := handler.New()

	log.Println("ping database")
	if err = storage.HealthCheck(); err != nil {
		log.Fatalln("ERROR: ", err)
	}

	log.Println("set server settings")
	srv := server.New(config.Cfg, handler)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Stop(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()
	log.Println("run server")
	if err = srv.Run(); err != http.ErrServerClosed {
		log.Fatalln("ERROR:", err)
	}
	<-idleConnsClosed
}
