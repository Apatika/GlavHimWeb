package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"log"
	"net/http"
	"strconv"
)

func getCities(w http.ResponseWriter, r *http.Request) {
	reg := r.URL.Query().Get("city")
	cities, err := service.GetCities(reg)
	if err != nil {
		errorResponse(w, "get cities failed", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(cities)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	reg := r.URL.Query().Get("customer")
	customers, err := service.GetCustomers(reg)
	if err != nil {
		errorResponse(w, fmt.Sprintf("get clients failed (%v)", err.Error()), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(customers)
}

func errorResponse(w http.ResponseWriter, errText string, code int) {
	log.Print(errText)
	http.Error(w, errText, code)
}

func searchOrders(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var payment bool
	if r.URL.Query().Get("payment") == "true" {
		payment = true
	} else if r.URL.Query().Get("payment") == "false" {
		payment = false
	} else {
		errorResponse(w, "payment error", http.StatusBadRequest)
		return
	}
	month := r.URL.Query().Get("month")
	limitInt, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		errorResponse(w, "convert limit to int failed", http.StatusBadRequest)
		return
	}
	limit := int64(limitInt)
	data, err := service.SearchOrders(id, payment, month, limit)
	if err != nil {
		errorResponse(w, "ошибка поиска", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}
