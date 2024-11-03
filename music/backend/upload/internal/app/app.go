package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"x-pltfrm/music/upload/config"
	"x-pltfrm/music/upload/internal/db"
	v1 "x-pltfrm/music/upload/internal/http/v1"
	"x-pltfrm/music/upload/pkg/logger"
)

func Run() {
	// Logger
	log := logger.New()
	log.Info("Logger initialized.")

	// Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}
	log.Info("Config initialized.", slog.Any("config", cfg))

	// DB connection
	db, err := db.NewPGPool(cfg.Postgres)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	log.Info("Database connected.")
	defer db.Close()

	// Server init
	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port),
		Handler: v1.Router(cfg.Routes),
	}

	// Startup
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting x-pltfrm/music/upload.")
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Server error.", err)
		}
	}()

	// Shutdown
	<-interrupt
	log.Info("Received interrupting signal. Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown failed.", err)
	}

	log.Info("Server stopped successfully.")
}
