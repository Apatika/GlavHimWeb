package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"html/template"
	"log"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /uri", uri)
	mux.HandleFunc("GET /db/{path}", dbPageGet)
	mux.HandleFunc("POST /db/{path}", dbPagePost)
	mux.HandleFunc("PUT /db/{path}", dbPagePut)
	mux.HandleFunc("DELETE /db/{path}", dbPageDelete)
	mux.HandleFunc("POST /orders", pushOrder)
	mux.HandleFunc("PUT /orders", updateOrder)
	mux.HandleFunc("PUT /orders/status", changeStatus)
	mux.HandleFunc("POST /clients", pushClient)
	mux.HandleFunc("PUT /clients", updateClient)
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

func uri(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		Uri string
	}{
		Uri: fmt.Sprintf("%v%v", config.Cfg.Server.URI, config.Cfg.Server.Port),
	})
}

func errorResponse(w http.ResponseWriter, errText string, code int) {
	log.Print(errText)
	http.Error(w, errText, code)
}
