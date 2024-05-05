package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
)

func getStructByPath(path string) service.IService {
	var data service.IService
	switch path {
	case "users":
		data = &service.User{}
	case "cargos":
		data = &service.Cargo{}
	case "chems":
		data = &service.Chemistry{}
	}
	return data
}

func dbPageGet(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	data, err := storage.GetAll(path)
	if err != nil {
		errorResponse(w, fmt.Sprintf("read db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func dbPagePost(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.NewID(storage.GetNewID())
	if err := storage.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error()), http.StatusBadRequest)
		return
	}
	if err := storage.AddOne(path, data); err != nil {
		errorResponse(w, fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("write to \"%v\" collection", path)
}

func dbPagePut(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := storage.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update db failed, path /db/%v (%v)", path, err.Error()), http.StatusBadRequest)
		return
	}
	if err := storage.UpdateOne(path, data, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("update to \"%v\" collection", path)
}

func dbPageDelete(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := storage.DeleteOne(path, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("delete from db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
		return
	}
	log.Printf("delete from \"%v\" collection", path)
}
