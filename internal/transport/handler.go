package handler

import (
	"net/http"
	"text/template"
)

const (
	def = 1 << iota
	checkName
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))

	mux.HandleFunc("/", index)

	mux.HandleFunc("GET /users", adminDBHandlerDefault)
	mux.HandleFunc("POST /users", adminDBHandlerUniqueName)
	mux.HandleFunc("PUT /users", adminDBHandlerUniqueName)
	mux.HandleFunc("DELETE /users", adminDBHandlerDefault)

	mux.HandleFunc("GET /cargos", adminDBHandlerDefault)
	mux.HandleFunc("POST /cargos", adminDBHandlerUniqueName)
	mux.HandleFunc("PUT /cargos", adminDBHandlerUniqueName)
	mux.HandleFunc("DELETE /cargos", adminDBHandlerDefault)

	mux.HandleFunc("GET /chems", adminDBHandlerDefault)
	mux.HandleFunc("POST /chems", adminDBHandlerUniqueName)
	mux.HandleFunc("PUT /chems", adminDBHandlerUniqueName)
	mux.HandleFunc("DELETE /chems", adminDBHandlerDefault)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}
