package handler

import (
	"encoding/json"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	StatusInWork  = "Принят В Работу"
	StatusShipped = "Отгружен"
	StatusPecom   = "Забор ПЭК"
	StatusCalled  = "Заказан Забор"
	StatusStop    = "СТОП"
	StatusCity    = "Развозка"
	StatusEmpty   = "Нет Товара"
)

func pushOrder(w http.ResponseWriter, r *http.Request) {
	path := "orders"
	var data service.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		return
	}
	if data.ID == primitive.NilObjectID {
		data.ID = primitive.NewObjectID()
	}
	var month time.Month
	data.CreationDate.Year, month, data.CreationDate.Day = time.Now().Date()
	data.CreationDate.Month = month.String()
	if err := storage.AddOne(path, data); err != nil {
		log.Printf("write db failed, path /db/%v (%v)", path, err.Error())
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
		return
	}
	json.NewEncoder(w).Encode(response{http.StatusOK, ""})
}

func inWorkOrders(w http.ResponseWriter, r *http.Request) {
	data, err := storage.GetInWorkOrders()
	if err != nil {
		log.Printf("read in work orders failed (%v)", err.Error())
		json.NewEncoder(w).Encode(response{http.StatusInternalServerError, err.Error()})
	}
	json.NewEncoder(w).Encode(data)
}
