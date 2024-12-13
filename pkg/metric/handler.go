package metric

import (
	"net/http"
	"todo/pkg/logger"

	"github.com/go-chi/chi"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
	// TODO: Logger
}

func (h *Handler) Register(router *chi.Mux) {
	router.Get(URL, h.Heartbeat)
	logger.Info("Registration handler heartbeat on /api/heartbeat")
}

func (h *Handler) Heartbeat(w http.ResponseWriter, r *http.Request) {
	logger.Info("[DEBUG] какой то ишак вызвал heartbeat:")
	w.WriteHeader(204)
}
