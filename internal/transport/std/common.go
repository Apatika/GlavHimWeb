package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"log"
	"net/http"
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
	reg := r.URL.Query().Get("client")
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
