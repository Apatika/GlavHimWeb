package handler

import (
	"encoding/json"
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

func adminDBHandler(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	data := getStructByPath(path)
	if r.Method != "GET" {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
			return
		}
	}
	switch r.Method {
	case http.MethodGet:
		data, err := storage.GetAll(path)
		if err != nil {
			log.Printf("read db failed, path /db/%v (%v)", path, err.Error())
			json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
			return
		}
		log.Printf("read from \"%v\" collection", path)
		json.NewEncoder(w).Encode(data)
		return
	case http.MethodPost:
		data.NewID(storage.GetNewID())
		if err := storage.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
			log.Printf("write db failed, path /db/%v (%v). %v", path, "name already exists", err.Error())
			json.NewEncoder(w).Encode(response{http.StatusBadRequest, err.Error()})
			return
		} else {
			if err := storage.AddOne(path, data); err != nil {
				log.Printf("write db failed, path /db/%v (%v). %v", path, err.Error(), err.Error())
				json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
				return
			}
		}
		log.Printf("write to \"%v\" collection", path)
	case http.MethodPut:
		if err := storage.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
			log.Printf("write db failed, path /db/%v (%v)", path, "name already exists")
			json.NewEncoder(w).Encode(response{http.StatusBadRequest, err.Error()})
			return
		} else {
			if err := storage.UpdateOne(path, data, data.GetID()); err != nil {
				log.Printf("write db failed, path /db/%v (%v)", path, err.Error())
				json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
				return
			}
		}
		log.Printf("update to \"%v\" collection", path)
	case http.MethodDelete:
		if err := storage.DeleteOne(path, data.GetID()); err != nil {
			log.Printf("delete from db failed, path /db/%v (%v)", path, err.Error())
			json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
			return
		}
		log.Printf("delete from \"%v\" collection", path)
	default:
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, "Method not alowed"})
		return
	}
	json.NewEncoder(w).Encode(response{http.StatusOK, ""})
}
