package handler

import (
	"encoding/json"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
	"strings"
)

func getStruct(path string) service.IService {
	switch path {
	case "users":
		return &service.User{}
	case "cargos":
		return &service.Cargo{}
	case "chems":
		return &service.Chemistry{}
	default:
		return nil
	}
}

func check(w http.ResponseWriter, r *http.Request, data service.IService, path string) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(data); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		return err
	}
	if err := storage.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusBadRequest, err.Error()})
		return err
	}
	return nil
}

func aDataGet(w http.ResponseWriter, r *http.Request) {
	path, found := strings.CutPrefix(r.URL.Path, "/")
	if !found {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, "url path: prefix not found"})
		log.Println("url path: prefix not found")
		return
	}
	data, err := storage.GetAll(path)
	if err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func aDataAdd(w http.ResponseWriter, r *http.Request) {
	path, found := strings.CutPrefix(r.URL.Path, "/")
	if !found {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, "url path: prefix not found"})
		log.Println("url path: prefix not found")
		return
	}
	data := getStruct(path)
	if err := check(w, r, data, path); err != nil {
		log.Println(err)
		return
	}
	data.NewID()
	if err := storage.AddOne(path, data); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(response{http.StatusOK, ""})
	log.Printf("added new value to \"%v\" collection: \"%v\" with id \"%v\"", path, data.GetName(), data.GetID())
}

func aDataUpdate(w http.ResponseWriter, r *http.Request) {
	path, found := strings.CutPrefix(r.URL.Path, "/")
	if !found {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, "url path: prefix not found"})
		log.Println("url path: prefix not found")
		return
	}
	data := getStruct(path)
	if err := check(w, r, data, path); err != nil {
		log.Println(err)
		return
	}
	if err := storage.UpdateOne(path, data, data.GetID()); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		log.Println(err)
	}
	json.NewEncoder(w).Encode(response{http.StatusOK, ""})
	log.Printf("update value at \"%v\" collection: \"%v\" with id \"%v\"", path, data.GetName(), data.GetID())
}

func aDataDelete(w http.ResponseWriter, r *http.Request) {
	path, found := strings.CutPrefix(r.URL.Path, "/")
	if !found {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, "url path: prefix not found"})
		log.Println("url path: prefix not found")
		return
	}
	data := getStruct(path)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(data); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		log.Println(err)
		return
	}
	if err := storage.DeleteOne(path, data.GetID()); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		log.Println(err)
	}
	json.NewEncoder(w).Encode(response{http.StatusOK, ""})
	log.Printf("delete value at \"%v\" collection: \"%v\" with id \"%v\"", path, data.GetName(), data.GetID())
}
