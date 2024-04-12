package handler

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getCargos(w http.ResponseWriter, r *http.Request) {
	cargos, err := storage.GetCargos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(cargos)
}

func checkCargoName(name string) error {
	cargos, err := storage.GetCargos()
	if err != nil {
		log.Println(err)
		return err
	}
	for _, value := range cargos {
		if value.Name == name {
			return fmt.Errorf("тк: имя занято")
		}
	}
	return nil
}

func addCargo(w http.ResponseWriter, r *http.Request) {
	var cargo service.Cargo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cargo); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := checkCargoName(cargo.Name); err != nil && cargo.ID == primitive.NilObjectID {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cargo.ID = primitive.NewObjectID()
	if err := storage.AddCargo(cargo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
	log.Printf("added new cargo: %v", cargo.Name)
}

func updateCargo(w http.ResponseWriter, r *http.Request) {
	var cargo service.Cargo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cargo); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := checkCargoName(cargo.Name); err != nil && cargo.ID == primitive.NilObjectID {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := storage.UpdateCargo(cargo); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("update cargo id: %v", cargo.ID)
}

func deleteCargo(w http.ResponseWriter, r *http.Request) {
	var cargo service.Cargo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cargo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err := storage.DeleteCargo(cargo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
	log.Printf("delete cargo id %v, name %v", cargo.ID, cargo.Name)
}
