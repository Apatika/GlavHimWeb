package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
)

func dbPageGet(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := service.GetStruct(path)
	json.NewEncoder(w).Encode(storage.Cache.Get(data.GetType()))
}

func dbPagePost(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := service.GetStruct(path)
	if data == nil {
		errorResponse(w, fmt.Sprintf("invalid path /db/%v", path), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	data.NewID(db.GetNewID())
	if err := db.CheckName(data.GetName(), path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("write to db failed(%v)", err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.Add(path, data); err != nil {
		errorResponse(w, fmt.Sprintf("write to db failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.Update(path, data.GetID(), data)
	log.Printf("write to \"%v\" collection", path)
}

func dbPagePut(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := service.GetStruct(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.CheckName(data.GetName(), path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update to db failed(%v)", err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.Update(path, data, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update to db failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.Update(path, data.GetID(), data)
	log.Printf("update to \"%v\" collection", path)
}

func dbPageDelete(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := service.GetStruct(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.Delete(path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("delete from db failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.Delete(path, data.GetID())
	log.Printf("delete from \"%v\" collection", path)
}
