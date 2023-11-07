package handler

import (
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {

}
