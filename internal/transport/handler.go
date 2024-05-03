package handler

import (
	"glavhim-app/internal/transport/std"
	"net/http"
)

func New() *http.ServeMux {
	return std.New()
}
