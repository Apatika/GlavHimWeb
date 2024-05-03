package std

import (
	"html/template"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /db/{path}", dbPageGet)
	mux.HandleFunc("POST /db/{path}", dbPagePost)
	mux.HandleFunc("PUT /db/{path}", dbPagePut)
	mux.HandleFunc("DELETE /db/{path}", dbPageDelete)
	mux.HandleFunc("POST /orders", pushOrder)
	mux.HandleFunc("PUT /orders/status", changeStatus)
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
