package handler

import (
	"encoding/json"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func pushClient(w http.ResponseWriter, r *http.Request) {
	path := "clients"
	var data service.Client
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		return
	}
	if data.ID == primitive.NilObjectID {
		data.ID = primitive.NewObjectID()
	}
	if err := storage.AddOne(path, data); err != nil {
		log.Printf("write db failed, path /db/%v (%v)", path, err.Error())
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		return
	}
	json.NewEncoder(w).Encode(struct {
		Code    int                `json:"code"`
		Message string             `json:"message"`
		ID      primitive.ObjectID `json:"id" bson:"_id"`
	}{http.StatusOK, "", data.ID})
}
