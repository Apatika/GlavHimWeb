package std

import (
	"encoding/json"
	"fmt"
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
		errText := fmt.Sprintf("decode client failed (%v)", err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	if data.ID == "" {
		data.ID = storage.GetNewID()
	}
	if err := storage.AddOne(path, data); err != nil {
		errText := fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	log.Printf("add new client (ID: %v)", data.ID)
	json.NewEncoder(w).Encode(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		ID      string `json:"id" bson:"_id"`
	}{http.StatusOK, "", data.ID})
}
