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

func dbPageGetManagers(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(storage.Cache.Managers())
}

func dbPagePostManagers(w http.ResponseWriter, r *http.Request) {
	var data service.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	data.NewID(db.GetNewID())
	if err := db.CheckNameOne(data.GetName(), config.Cfg.DB.Coll.Users, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("write user to db failed(%v)", err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.AddOne(config.Cfg.DB.Coll.Users, data); err != nil {
		errorResponse(w, fmt.Sprintf("write user to db failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.AddToManagers(data)
	log.Printf("write to \"%v\" collection", config.Cfg.DB.Coll.Users)
}

func dbPagePutManagers(w http.ResponseWriter, r *http.Request) {
	var data service.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.CheckNameOne(data.GetName(), config.Cfg.DB.Coll.Users, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update user to db failed(%v)", err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.UpdateOne(config.Cfg.DB.Coll.Users, data, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update user to db failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.AddToManagers(data)
	log.Printf("update to \"%v\" collection", config.Cfg.DB.Coll.Users)
}

func dbPageDeleteManagers(w http.ResponseWriter, r *http.Request) {
	var data service.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.DeleteOne(config.Cfg.DB.Coll.Users, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("delete user from db failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.DeleteManagers(data)
	log.Printf("delete from \"%v\" collection", config.Cfg.DB.Coll.Users)
}
