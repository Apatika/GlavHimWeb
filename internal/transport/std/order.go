package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/config"
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
	path := config.Cfg.DB.Coll.Orders
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
			if err := db.AddOne(config.Cfg.DB.Coll.Clients, data.Client); err != nil {
				errorResponse(w, fmt.Sprintf("push client failed(%v)", err.Error()), http.StatusInternalServerError)
				return
			}
		} else {
			data.Client.ID = id
			if err := db.UpdateOne(config.Cfg.DB.Coll.Clients, data.Client, data.Client.ID); err != nil {
				errorResponse(w, fmt.Sprintf("update client failed(%v)", err.Error()), http.StatusInternalServerError)
				return
			}
		}
	} else {
		if err := db.UpdateOne(config.Cfg.DB.Coll.Clients, data.Client, data.Client.ID); err != nil {
			errorResponse(w, fmt.Sprintf("update client failed(%v)", err.Error()), http.StatusInternalServerError)
			return
		}
	}
	data.Order.ClientID = data.Client.ID
	if err := db.AddOne(path, data.Order); err != nil {
		errorResponse(w, fmt.Sprintf("push order failed(%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.Update(path, data.Order.ID, data)
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
	if err := db.UpdateOne(config.Cfg.DB.Coll.Orders, data.Order, data.Order.ID); err != nil {
		errorResponse(w, fmt.Sprintf("order update failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := db.UpdateOne(config.Cfg.DB.Coll.Clients, data.Client, data.Client.ID); err != nil {
		errorResponse(w, fmt.Sprintf("client update failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.Update(config.Cfg.DB.Coll.Orders, data.Order.ID, data)
	log.Printf("update order ID: %v", data.Order.ID)
}

func inWorkOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.Cache.Get(config.Cfg.DB.Coll.Orders))
}

func changeStatus(w http.ResponseWriter, r *http.Request) {
	var data service.OrderFull
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		errorResponse(w, fmt.Sprintf("decode newStatus failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if data.Order.ID == "" {
		errorResponse(w, "invalid field values (%v)", http.StatusBadRequest)
		return
	}
	if data.Order.Status == StatusShipped {
		var month time.Month
		data.Order.ShipmentDate.Year, month, data.Order.ShipmentDate.Day = time.Now().Date()
		data.Order.ShipmentDate.Month = month.String()
	}
	db := storage.DB()
	if err := db.UpdateOne(config.Cfg.DB.Coll.Orders, data.Order, data.Order.ID); err != nil {
		errorResponse(w, fmt.Sprintf("update status failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	storage.Cache.Update(config.Cfg.DB.Coll.Orders, data.Order.ID, data)
	log.Printf("update status ID: %v, status: %v", data.Order.ID, data.Order.Status)
	json.NewEncoder(w).Encode(storage.Cache.Get(config.Cfg.DB.Coll.Orders))
}
