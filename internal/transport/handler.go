package handler

import (
	"encoding/json"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /users", getAll)
	mux.HandleFunc("POST /users", addOne)
	mux.HandleFunc("PUT /users", updateOne)
	mux.HandleFunc("DELETE /users", deleteOne)
	mux.HandleFunc("GET /cargos", getAll)
	mux.HandleFunc("POST /cargos", addOne)
	mux.HandleFunc("PUT /cargos", updateOne)
	mux.HandleFunc("DELETE /cargos", deleteOne)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func getStruct(path string) service.IService {
	switch path {
	case "users":
		return &service.User{}
	case "cargos":
		return &service.Cargo{}
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

func getAll(w http.ResponseWriter, r *http.Request) {
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

func addOne(w http.ResponseWriter, r *http.Request) {
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

func updateOne(w http.ResponseWriter, r *http.Request) {
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

func deleteOne(w http.ResponseWriter, r *http.Request) {
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
