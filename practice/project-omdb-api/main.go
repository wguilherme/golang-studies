package main

import (
	"log/slog"
	"net/http"
	"os"
	"project-url-shortener/api"
	"time"
)

func main() {

	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")

}

func run() error {

	apiKey := os.Getenv("OMDB_API_KEY")

	handler := api.NewHandler(apiKey)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
