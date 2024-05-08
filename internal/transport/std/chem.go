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

func dbPageGetChemistry(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(storage.Cache.Chemistry())
}

func dbPagePostChemistry(w http.ResponseWriter, r *http.Request) {
	var data service.Chemistry
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	data.NewID(db.GetNewID())
	if err := db.CheckNameOne(data.GetName(), config.Cfg.DB.Coll.Chems, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("write chemistry to db failed(%v)", err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.AddOne(config.Cfg.DB.Coll.Chems, data); err != nil {
		errorResponse(w, fmt.Sprintf("write chemistry to db failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.AddToChemistry(data)
	log.Printf("write to \"%v\" collection", config.Cfg.DB.Coll.Chems)
}

func dbPagePutChemistry(w http.ResponseWriter, r *http.Request) {
	var data service.Chemistry
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.CheckNameOne(data.GetName(), config.Cfg.DB.Coll.Chems, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update chemistry to db failed(%v)", err.Error()), http.StatusBadRequest)
		return
	}
	if err := db.UpdateOne(config.Cfg.DB.Coll.Chems, data, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("update chemistry to db failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.AddToChemistry(data)
	log.Printf("update to \"%v\" collection", config.Cfg.DB.Coll.Chems)
}

func dbPageDeleteChemistry(w http.ResponseWriter, r *http.Request) {
	var data service.Chemistry
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.DeleteOne(config.Cfg.DB.Coll.Chems, data.GetID()); err != nil {
		errorResponse(w, fmt.Sprintf("delete from db failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.DeleteChemistry(data)
	log.Printf("delete from \"%v\" collection", config.Cfg.DB.Coll.Chems)
}
