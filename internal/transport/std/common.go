package std

import (
	"encoding/json"
	"glavhim-app/internal/storage"
	"net/http"
)

func getCities(w http.ResponseWriter, r *http.Request) {
	reg := r.URL.Query().Get("city")
	db := storage.DB()
	result, err := db.GetCities(reg)
	if err != nil {
		errorResponse(w, "get cities failed", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}
