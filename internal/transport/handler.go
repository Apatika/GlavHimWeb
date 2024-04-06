package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		fmt.Println("index.html parse error")
	}
	t.ExecuteTemplate(w, "index.html", nil)
}
