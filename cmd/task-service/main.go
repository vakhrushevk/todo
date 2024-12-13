package main

import (
	"log/slog"
	"todo/internal/config"
	"todo/pkg/logger"
	"todo/pkg/metric"

	"github.com/go-chi/chi"
)

func main() {
	//logger.Init()
	logger.Info("Логер инициализирован")

	cfg := config.GetConfig()

	router := chi.NewMux()
	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	// TODO: db
	// TODO: repository
	// TODO: service
	// TODO: handlers
	// TODO: start

	logger.Info("start http server on ", slog.String("host", cfg.Http.Port), slog.String("port", cfg.Http.Port))

	//	http.ListenAndServe(cfg.Http.Address(), router)
}
