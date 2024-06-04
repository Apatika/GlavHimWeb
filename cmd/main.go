package main

import (
	"context"
	"glavhim-app/internal/config"
	"glavhim-app/internal/server"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	handler "glavhim-app/internal/transport"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var DB_URI string

func main() {

	log.SetFlags(log.LstdFlags)
	log.Print("open log file")
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalf("ERROR: open log file failed (%v)", err)
	}
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	config.New()

	log.Println("load http handler")
	handler := handler.New()

	storage.NewDB()

	storage.CacheInit()
	if err := service.InWorkToCache(); err != nil {
		log.Fatalf("ERROR: load cache failed (%v)", err)
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
	log.Println("start server")
	if err = srv.Run(); err != http.ErrServerClosed {
		log.Fatalln("ERROR:", err)
	}
	<-idleConnsClosed
}
