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
	data.Order.ID = storage.GetNewID()
	var month time.Month
	data.Order.CreationDate.Year, month, data.Order.CreationDate.Day = time.Now().Date()
	data.Order.CreationDate.Month = month.String()
	if data.Client.ID == "" {
		id, err := storage.CheckClient(data.Client)
		if err != nil {
			data.Client.ID = storage.GetNewID()
			if err := storage.AddOne("clients", data.Client); err != nil {
				errorResponse(w, fmt.Sprintf("push client failed(%v)", err.Error()), http.StatusInternalServerError)
				return
			}
		} else {
			data.Client.ID = id
			if err := storage.UpdateOne("clients", data.Client, data.Client.ID); err != nil {
				errorResponse(w, fmt.Sprintf("update client failed(%v)", err.Error()), http.StatusInternalServerError)
				return
			}
		}
	} else {
		if err := storage.UpdateOne("clients", data.Client, data.Client.ID); err != nil {
			errorResponse(w, fmt.Sprintf("update client failed(%v)", err.Error()), http.StatusInternalServerError)
			return
		}
	}
	data.Order.ClientID = data.Client.ID
	if err := storage.AddOne(path, data.Order); err != nil {
		errorResponse(w, fmt.Sprintf("push order failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := storage.CacheUpdateInWork(); err != nil {
		log.Print("cache update failed")
	}
	log.Printf("add new order, invoice %v (ID: %v)", data.Order.Invoice, data.Order.ID)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	path := "orders"
	var data service.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, fmt.Sprintf("decode order failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := storage.UpdateOne(path, data, data.ID); err != nil {
		errorResponse(w, fmt.Sprintf("write db failed, path /db/%v (%v)", path, err.Error()), http.StatusInternalServerError)
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
		ID           string       `json:"id" bson:"_id"`
		Status       string       `json:"status" bson:"status"`
		ShipmentDate service.Date `json:"shipmentDate" bson:"shipment_date"`
	}{}
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
	if err := storage.UpdateOne("orders", newStatus, newStatus.ID); err != nil {
		errorResponse(w, fmt.Sprintf("update status failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := storage.CacheUpdateInWork(); err != nil {
		log.Print("cache update failed")
	}
	log.Printf("update status ID: %v, status: %v", newStatus.ID, newStatus.Status)
}
