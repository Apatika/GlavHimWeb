package handler

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
	"strings"
)

func adminDBHandlerDefault(w http.ResponseWriter, r *http.Request) {
	code, err := adminDBHandler(w, r, def)
	if r.Method != "GET" {
		json.NewEncoder(w).Encode(response{code, fmt.Sprintf("%v (code %v)", err, code)})
	}
	if code != http.StatusOK {
		log.Println(err)
	}
}

func adminDBHandlerUniqueName(w http.ResponseWriter, r *http.Request) {
	code, err := adminDBHandler(w, r, checkName)
	json.NewEncoder(w).Encode(response{code, fmt.Sprintf("%v (code %v)", err, code)})
	if code != http.StatusOK {
		log.Println(err)
	}
}

func adminDBHandler(w http.ResponseWriter, r *http.Request, flag int) (int, string) {
	path, found := strings.CutPrefix(r.URL.Path, "/")
	if !found {
		return http.StatusInternalServerError, fmt.Sprintf("url path: prefix %v not found", path)
	}
	var data service.IService
	if r.Method != "GET" {
		switch path {
		case "users":
			data = &service.User{}
		case "cargos":
			data = &service.Cargo{}
		case "chems":
			data = &service.Chemistry{}
		}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		if flag&checkName == checkName {
			if err := storage.CheckNameOne(data.GetName(), path, data.GetID()); err != nil {
				return http.StatusBadRequest, err.Error()
			}
		}
	}
	switch r.Method {
	case http.MethodGet:
		data, err := storage.GetAll(path)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		json.NewEncoder(w).Encode(data)
	case http.MethodPost:
		data.NewID()
		if err := storage.AddOne(path, data); err != nil {
			return http.StatusInternalServerError, err.Error()
		}
	case http.MethodPut:
		if err := storage.UpdateOne(path, data, data.GetID()); err != nil {
			return http.StatusInternalServerError, err.Error()
		}
	case http.MethodDelete:
		if err := storage.DeleteOne(path, data.GetID()); err != nil {
			return http.StatusInternalServerError, err.Error()
		}
	default:
		return http.StatusInternalServerError, "Method not allowed"
	}
	return http.StatusOK, ""
}
