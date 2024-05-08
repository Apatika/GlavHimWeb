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

func getStructByPath(path string) service.IDBPage {
	var data service.IDBPage
	switch path {
	case config.Cfg.DB.Coll.Users:
		data = &service.User{}
	case config.Cfg.DB.Coll.Cargos:
		data = &service.Cargo{}
	case config.Cfg.DB.Coll.Chems:
		data = &service.Chemistry{}
	}
	return data
}

func dbPageGet(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	cache, err := storage.Cache()
	if err != nil {
		errorResponse(w, fmt.Sprintf("cache read failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	switch path {
	case config.Cfg.DB.Coll.Users:
		json.NewEncoder(w).Encode(cache.Managers())
	case config.Cfg.DB.Coll.Cargos:
		json.NewEncoder(w).Encode(cache.Cargos())
	case config.Cfg.DB.Coll.Chems:
		json.NewEncoder(w).Encode(cache.Chemistry())
	default:
		errorResponse(w, fmt.Sprintf("invalid path /db/%v", path), http.StatusInternalServerError)
	}
}

func dbPagePost(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	data.NewID(db.GetNewID())
	if err := db.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.AddOne(path, data); err != nil {
		errorResponse(w, fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("write to \"%v\" collection", path)
}

func dbPagePut(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update db failed, path /db/%v (%v)", path, err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.UpdateOne(path, data, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("update to \"%v\" collection", path)
}

func dbPageDelete(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.DeleteOne(path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("delete from db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("delete from \"%v\" collection", path)
}
