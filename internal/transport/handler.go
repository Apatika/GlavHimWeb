package handler

import (
	"net/http"
	"text/template"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/db/{path}", adminDBHandler)
	mux.HandleFunc("POST /orders", pushOrder)
	mux.HandleFunc("POST /clients", pushClient)
	mux.HandleFunc("GET /inwork", inWorkOrders)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}
