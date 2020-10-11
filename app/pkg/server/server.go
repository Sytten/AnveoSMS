package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/sytten/anveosms/pkg/sms"
)

type Server struct {
	sms sms.Service

	logger *zap.Logger

	router chi.Router
}

func New(sms sms.Service, logger *zap.Logger) *Server {
	r := chi.NewRouter()

	r.Route("/webhooks", func(r chi.Router) {
		h := webhooksHandler{sms}
		r.Mount("/", h.route())
	})

	return &Server{sms, logger, r}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
