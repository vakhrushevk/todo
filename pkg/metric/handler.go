package metric

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
	// TODO: Logger
}

func (h *Handler) Register(router *chi.Mux) {
	router.Get(URL, h.Heartbeat)
	log.Printf("[I]NFO] Registration handler Heartbeat on /api/heartbeat")
	//TODO: LOGGER

}

func (h *Handler) Heartbeat(w http.ResponseWriter, r *http.Request) {
	log.Printf("[DEBUG] какой то ишак вызвал heartbeat:")
	w.WriteHeader(204)
}

echo "# todo" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/vakhrushevk/todo.git
git push -u origin main