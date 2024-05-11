package std

import (
	"encoding/json"
	"glavhim-app/internal/service"
	"io"
	"net/http"
)

func dbPageGet(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := service.GetStruct(path).GetAll()
	json.NewEncoder(w).Encode(data)
}

func dbPagePost(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := service.GetStruct(path)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := data.Push(body); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func dbPagePut(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := service.GetStruct(path)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := data.Update(body); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func dbPageDelete(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	data := service.GetStruct(path)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := data.Delete(body); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
