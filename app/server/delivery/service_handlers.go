package delivery

import (
	"net/http"
)

func (h *Handler) Clear(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	code := 200

	err := h.Service.Clear()
	if err != nil {
		code = 500
	}

	w.WriteHeader(code)
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
