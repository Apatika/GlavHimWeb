package handler

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
	"text/template"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("POST /users", addUser)
	mux.HandleFunc("PUT /users", updateUser)
	mux.HandleFunc("DELETE /users", deleteUser)
	mux.HandleFunc("GET /cargos", getCargos)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		fmt.Println("index.html parse error")
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func getCargos(w http.ResponseWriter, r *http.Request) {
	cargos, err := storage.GetCargos()
	if err != nil {
		log.Println("getting cargos error")
		return
	}
	json.NewEncoder(w).Encode(cargos)
}
