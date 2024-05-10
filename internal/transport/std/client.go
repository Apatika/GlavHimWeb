package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
)

func updateClient(w http.ResponseWriter, r *http.Request) {
	path := config.Cfg.DB.Coll.Clients
	var data service.Client
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, fmt.Sprintf("decode client failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.Update(path, data, data.ID); err != nil {
		errorResponse(w, fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("update client ID: %v", data.ID)
}
