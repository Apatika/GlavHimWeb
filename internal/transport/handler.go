package handler

import (
	"net/http"
	"text/template"
)

const (
	def = 1 << iota
	checkName
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))

	mux.HandleFunc("/", index)

	mux.HandleFunc("GET /db/{path}", adminDBHandlerDefault)
	mux.HandleFunc("POST /db/{path}", adminDBHandlerUniqueName)
	mux.HandleFunc("PUT /db/{path}", adminDBHandlerUniqueName)
	mux.HandleFunc("DELETE /db/{path}", adminDBHandlerDefault)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}
