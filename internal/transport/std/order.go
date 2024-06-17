package std

import (
	"fmt"
	"glavhim-app/internal/service"
	"io"
	"net/http"
)

func orderPush(w http.ResponseWriter, r *http.Request) {
	defer clients.Update()
	var order service.Order
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

func orderUpdate(w http.ResponseWriter, r *http.Request) {
	defer clients.Update()
	var order service.Order
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

func orderDelete(w http.ResponseWriter, r *http.Request) {
	defer clients.Update()
	var order service.Order
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, fmt.Sprintf("decode order failed (%v)", err.Error()), http.StatusInternalServerError)
		return
	}
	if err := order.Delete(body); err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
