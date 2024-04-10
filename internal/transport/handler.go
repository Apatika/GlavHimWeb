package handler

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/storage"
	"net/http"
	"text/template"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /users", users)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		fmt.Println("index.html parse error")
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsers()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(users)
}
