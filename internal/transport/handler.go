package handler

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
	"text/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("POST /users", addUser)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		fmt.Println("index.html parse error")
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsers()
	if err != nil {
		log.Println("getting users error")
		return
	}
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user service.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println("add user error")
		return
	}
	user.ID = primitive.NewObjectID()
	storage.AddUser(user)
	log.Printf("added new user: %v", user.Name)
}
