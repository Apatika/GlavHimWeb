package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"html/template"
	"log"
	"net/http"
)

const pathVar = "path"

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /settings", settings)
	mux.HandleFunc(fmt.Sprintf("GET /db/{%v}", pathVar), dbPageGet)
	mux.HandleFunc(fmt.Sprintf("POST /db/{%v}", pathVar), dbPagePost)
	mux.HandleFunc(fmt.Sprintf("PUT /db/{%v}", pathVar), dbPagePut)
	mux.HandleFunc(fmt.Sprintf("DELETE /db/{%v}", pathVar), dbPageDelete)
	mux.HandleFunc("POST /orders", pushOrder)
	mux.HandleFunc("PUT /orders", updateOrder)
	mux.HandleFunc("PUT /orders/status", updateOrder)
	mux.HandleFunc("GET /inwork", inWorkOrders)
	mux.HandleFunc("GET /cities", getCities)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func settings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		RefreshRate int `json:"refreshRate"`
	}{
		RefreshRate: config.Cfg.Frontend.RefreshRate,
	})
}

func errorResponse(w http.ResponseWriter, errText string, code int) {
	log.Print(errText)
	http.Error(w, errText, code)
}
