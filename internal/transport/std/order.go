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
	StatusChanged = "Изменен!"
)

func pushOrder(w http.ResponseWriter, r *http.Request) {
	path := "orders"
	var data service.OrderFull
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, fmt.Sprintf("decode order failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	data.Order.ID = db.GetNewID()
	var month time.Month
	data.Order.CreationDate.Year, month, data.Order.CreationDate.Day = time.Now().Date()
	data.Order.CreationDate.Month = month.String()
	if data.Client.ID == "" {
		id, err := db.CheckClient(data.Client)
		if err != nil {
			data.Client.ID = db.GetNewID()
			if err := db.AddOne("clients", data.Client); err != nil {
				errorResponse(w, fmt.Sprintf("push client failed(%v)", err.Error()), http.StatusInternalServerError)
				return
			}
		} else {
			data.Client.ID = id
			if err := db.UpdateOne("clients", data.Client, data.Client.ID); err != nil {
				errorResponse(w, fmt.Sprintf("update client failed(%v)", err.Error()), http.StatusInternalServerError)
				return
			}
		}
	} else {
		if err := db.UpdateOne("clients", data.Client, data.Client.ID); err != nil {
			errorResponse(w, fmt.Sprintf("update client failed(%v)", err.Error()), http.StatusInternalServerError)
			return
		}
	}
	data.Order.ClientID = data.Client.ID
	if err := db.AddOne(path, data.Order); err != nil {
		errorResponse(w, fmt.Sprintf("push order failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.CacheNewOrder(data)
	log.Printf("add new order, invoice %v (ID: %v)", data.Order.Invoice, data.Order.ID)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	var data service.OrderFull
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, fmt.Sprintf("decode order failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	db := storage.DB()
	if err := db.UpdateOne("orders", data.Order, data.Order.ID); err != nil {
		errorResponse(w, fmt.Sprintf("order update failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := db.UpdateOne("clients", data.Client, data.Client.ID); err != nil {
		errorResponse(w, fmt.Sprintf("client update failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := storage.CacheUpdateOrder(data); err != nil {
		log.Print(err)
	}
	log.Printf("update order ID: %v", data.Order.ID)
}

func inWorkOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.CacheGetInWork())
}

func changeStatus(w http.ResponseWriter, r *http.Request) {
	var newStatus service.OrderStatusChanger
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newStatus); err != nil {
		errorResponse(w, fmt.Sprintf("decode newStatus failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if newStatus.ID == "" {
		errorResponse(w, "invalid field values (%v)", http.StatusBadRequest)
		return
	}
	if newStatus.Status == StatusShipped {
		var month time.Month
		newStatus.ShipmentDate.Year, month, newStatus.ShipmentDate.Day = time.Now().Date()
		newStatus.ShipmentDate.Month = month.String()
	}
	db := storage.DB()
	if err := db.UpdateOne("orders", newStatus, newStatus.ID); err != nil {
		errorResponse(w, fmt.Sprintf("update status failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := storage.CacheUpdateOrderStatus(newStatus); err != nil {
		log.Print("cache update failed")
	}
	log.Printf("update status ID: %v, status: %v", newStatus.ID, newStatus.Status)
	json.NewEncoder(w).Encode(storage.CacheGetInWork())
}
