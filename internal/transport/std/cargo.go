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

func dbPageGetCargo(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(storage.Cache.Cargos())
}

func dbPagePostCargo(w http.ResponseWriter, r *http.Request) {
	var data service.Cargo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	data.NewID(db.GetNewID())
	if err := db.CheckNameOne(data.GetName(), config.Cfg.DB.Coll.Cargos, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("write cargo to db failed(%v)", err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.AddOne(config.Cfg.DB.Coll.Cargos, data); err != nil {
		errorResponse(w, fmt.Sprintf("write cargo to db failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.AddToCargos(data)
	log.Printf("write to \"%v\" collection", config.Cfg.DB.Coll.Cargos)
}

func dbPagePutCargo(w http.ResponseWriter, r *http.Request) {
	var data service.Cargo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.CheckNameOne(data.GetName(), config.Cfg.DB.Coll.Cargos, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update cargo to db failed(%v)", err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.UpdateOne(config.Cfg.DB.Coll.Cargos, data, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update cargo to db failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.AddToCargos(data)
	log.Printf("update to \"%v\" collection", config.Cfg.DB.Coll.Cargos)
}

func dbPageDeleteCargo(w http.ResponseWriter, r *http.Request) {
	var data service.Cargo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.DeleteOne(config.Cfg.DB.Coll.Cargos, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("delete cargo from db failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.DeleteCargos(data)
	log.Printf("delete from \"%v\" collection", config.Cfg.DB.Coll.Cargos)
}
