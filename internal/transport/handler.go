package handler

import (
	"net/http"
	"text/template"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))

	mux.HandleFunc("/", index)

	mux.HandleFunc("GET /users", aDataGet)
	mux.HandleFunc("POST /users", aDataAdd)
	mux.HandleFunc("PUT /users", aDataUpdate)
	mux.HandleFunc("DELETE /users", aDataDelete)

	mux.HandleFunc("GET /cargos", aDataGet)
	mux.HandleFunc("POST /cargos", aDataAdd)
	mux.HandleFunc("PUT /cargos", aDataUpdate)
	mux.HandleFunc("DELETE /cargos", aDataDelete)

	mux.HandleFunc("GET /chems", aDataGet)
	mux.HandleFunc("POST /chems", aDataAdd)
	mux.HandleFunc("PUT /chems", aDataUpdate)
	mux.HandleFunc("DELETE /chems", aDataDelete)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}
