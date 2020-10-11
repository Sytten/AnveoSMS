package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/sytten/anveosms/pkg/sms"
)

type webhooksHandler struct {
	s sms.Service
}

func (h *webhooksHandler) route() chi.Router {
	r := chi.NewRouter()

	r.Post("/anveo", h.receiveAnveoSMS)

	return r
}

func (h *webhooksHandler) receiveAnveoSMS(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	message := r.URL.Query().Get("message")

	h.s.Receive(from, to, message)

	render.Status(r, http.StatusOK)
}
