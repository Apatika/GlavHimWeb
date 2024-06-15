package std

import (
	"encoding/json"
	"fmt"
	"glavhim-app/internal/service"
	"io"
	"net/http"
)

func getPvd(w http.ResponseWriter, r *http.Request) {
	pvds, err := service.GetPvds()
	if err != nil {
		errorResponse(w, fmt.Sprintf("get pvds failed (%v)", err.Error()), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(pvds)
}

func addPvd(w http.ResponseWriter, r *http.Request) {
	var pvd service.Pvd
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, fmt.Sprintf("decode pvd failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := pvd.Push(body); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := service.GetPvds()
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}
