package handler

import (
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
	mux.HandleFunc("POST /cargos", addCargo)
	mux.HandleFunc("PUT /cargos", updateCargo)
	mux.HandleFunc("DELETE /cargos", deleteCargo)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}
