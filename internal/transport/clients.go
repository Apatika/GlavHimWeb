package handler

import (
	"encoding/json"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
)

func pushClient(w http.ResponseWriter, r *http.Request) {
	path := "clients"
	var data service.Client
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		return
	}
	if data.ID == "" {
		data.ID = storage.GetNewID()
	}
	if err := storage.AddOne(path, data); err != nil {
		log.Printf("write db failed, path /db/%v (%v)", path, err.Error())
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		return
	}
	json.NewEncoder(w).Encode(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		ID      string `json:"id" bson:"_id"`
	}{http.StatusOK, "", data.ID})
}
