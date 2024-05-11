package std

import (
	"encoding/json"
	"glavhim-app/internal/service"
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
