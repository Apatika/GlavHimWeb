package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"glavhim-app/internal/storage"
	"log"
	"net/http"
	"time"
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
		errText := fmt.Sprintf("decode order failed (%v)", err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	if data.ID == "" {
		data.ID = storage.GetNewID()
	}
	var month time.Month
	data.CreationDate.Year, month, data.CreationDate.Day = time.Now().Date()
	data.CreationDate.Month = month.String()
	if err := storage.AddOne(path, data); err != nil {
		errText := fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	if err := storage.CacheUpdateInWork(); err != nil {
		log.Print("cache update failed")
	}
	log.Printf("add new order, invoice %v (ID: %v)", data.Invoice, data.ID)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	path := "orders"
	var data service.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errText := fmt.Sprintf("decode order failed (%v)", err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	if err := storage.UpdateOne(path, data, data.ID); err != nil {
		errText := fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	if err := storage.CacheUpdateInWork(); err != nil {
		log.Print("cache update failed")
	}
	log.Printf("update order ID: %v", data.ID)
}

func inWorkOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.CacheGetInWork())
}

func changeStatus(w http.ResponseWriter, r *http.Request) {
	newStatus := struct {
		ID     string `json:"id" bson:"_id"`
		Status string `json:"status" bson:"status"`
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newStatus); err != nil {
		errText := fmt.Sprintf("decode newStatus failed (%v)", err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	if newStatus.ID == "" || newStatus.Status == "" {
		errText := "invalid field values (%v)"
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	if err := storage.UpdateOne("orders", newStatus, newStatus.ID); err != nil {
		errText := fmt.Sprintf("update status failed (%v)", err.Error())
		log.Print(errText)
		http.Error(w, errText, http.StatusInternalServerError)
		return
	}
	if err := storage.CacheUpdateInWork(); err != nil {
		log.Print("cache update failed")
	}
	log.Printf("update status ID: %v, status: %v", newStatus.ID, newStatus.Status)
}
