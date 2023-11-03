package gorillamux

import (
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {

}

func New() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", index)

	return r
}
