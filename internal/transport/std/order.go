package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"io"
	"net/http"
)

func pushOrder(w http.ResponseWriter, r *http.Request) {
	order := service.NewOrderFull()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, fmt.Sprintf("decode order failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := order.Push(body); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	order := service.NewOrderFull()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, fmt.Sprintf("decode order failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := order.Update(body); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func inWorkOrders(w http.ResponseWriter, r *http.Request) {
	order := service.NewOrderFull()
	data := order.GetAll()
	json.NewEncoder(w).Encode(data)
}
