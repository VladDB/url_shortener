package main

import (
	"log/slog"
	"os"
	"ulr_shortener/internal/config"
	"ulr_shortener/internal/lib/logger/sl"
	"ulr_shortener/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// read config
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug message are enabled")

	// TODO: init storage: sqlite
	storage, err := postgres.NewStorage()
	if err != nil {
		log.Error("filed to init storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage
	// TODO: init router:  chi "chi render"

	// TODO: init server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
