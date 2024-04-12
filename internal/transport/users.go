package handler

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("getting users error")
		return
	}
	json.NewEncoder(w).Encode(users)
}

func checkUserName(name string) error {
	users, err := storage.GetUsers()
	if err != nil {
		log.Println(err)
		return err
	}
	for _, value := range users {
		if value.Name == name {
			return fmt.Errorf("менеджеры: имя занято")
		}
	}
	return nil
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user service.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := checkUserName(user.Name); err != nil && user.ID == primitive.NilObjectID {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	user.ID = primitive.NewObjectID()
	if err := storage.AddUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
	log.Printf("added new user: %v", user.Name)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user service.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := checkUserName(user.Name); err != nil && user.ID == primitive.NilObjectID {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	if err := storage.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
	log.Printf("update user id: %v", user.ID)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var user service.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := storage.DeleteUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
	log.Printf("delete user id %v, name %v", user.ID, user.Name)
}
