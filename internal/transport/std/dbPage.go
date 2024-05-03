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
		errText := fmt.Sprintf("read db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func dbPagePost(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.NewID(storage.GetNewID())
	if err := storage.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
		errText := fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusBadRequest)
		return
	}
	if err := storage.AddOne(path, data); err != nil {
		errText := fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	log.Printf("write to \"%v\" collection", path)
}

func dbPagePut(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := storage.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
		errText := fmt.Sprintf("update db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusBadRequest)
		return
	}
	if err := storage.UpdateOne(path, data, data.GetID()); err != nil {
		errText := fmt.Sprintf("update db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	log.Printf("update to \"%v\" collection", path)
}

func dbPageDelete(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	data := getStructByPath(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := storage.DeleteOne(path, data.GetID()); err != nil {
		errText := fmt.Sprintf("delete from db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	log.Printf("delete from \"%v\" collection", path)
}
