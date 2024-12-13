package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"todo/internal/config"
	"todo/pkg/metric"
)

func main() {

	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// TODO: logger
	cfg := config.GetConfig()

	router := chi.NewMux()

	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	// TODO: db
	// TODO: repository
	// TODO: service
	// TODO: handlers
	// TODO: start

	fmt.Printf("Start http server on %v", cfg.Http.Address())

	http.ListenAndServe(cfg.Http.Address(), router)
}
