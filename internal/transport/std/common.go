package std

import (
	"encoding/json"
	"glavhim-app/internal/config"
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

func settings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		RefreshRate int `json:"refreshRate"`
	}{
		RefreshRate: config.Cfg.Frontend.RefreshRate,
	})
}

func errorResponse(w http.ResponseWriter, errText string, code int) {
	log.Print(errText)
	http.Error(w, errText, code)
}
