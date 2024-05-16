package std

import (
	"fmt"
	"glavhim-app/internal/service"
	"html/template"
	"net/http"

	"github.com/gorilla/websocket"
)

const pathVar = "path"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc(fmt.Sprintf("GET /db/{%v}", pathVar), dbPageGet)
	mux.HandleFunc(fmt.Sprintf("POST /db/{%v}", pathVar), dbPagePost)
	mux.HandleFunc(fmt.Sprintf("PUT /db/{%v}", pathVar), dbPagePut)
	mux.HandleFunc(fmt.Sprintf("DELETE /db/{%v}", pathVar), dbPageDelete)
	mux.HandleFunc("POST /orders", pushOrder)
	mux.HandleFunc("PUT /orders", updateOrder)
	mux.HandleFunc("PUT /orders/status", updateOrder)
	mux.HandleFunc("GET /inwork", wsHandler)
	mux.HandleFunc("GET /cities", getCities)
	mux.HandleFunc("GET /clients", getClients)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	ch := make(chan struct{}, 1)
	id := clients.AddChan(ch)
	ch <- struct{}{}
	for {
		<-ch
		order := service.NewOrderFull()
		data := order.GetAll()
		if err := conn.WriteJSON(data); err != nil {
			clients.DeleteChan(id)
			conn.Close()
		}
	}
}
