package router

import (
	gorillamux "apatikaWebServer/internal/transport/rest/gorillaMux"
	"log"
	"net/http"
	"os"
)

func New() http.Handler {
	var handler http.Handler
	pkgName := os.Getenv("MUX")
	switch pkgName {
	case "gorilla":
		handler = gorillamux.New()
	default:
		log.Fatal("Not allowed http handler")
	}
	return handler
}
