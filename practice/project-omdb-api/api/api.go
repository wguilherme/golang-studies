package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"project-url-shortener/omdb"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewHandler(apiKey string) http.Handler {
	// inicializa o router
	r := chi.NewMux()

	// aplica middlewares
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// declara as rotas
	r.Get("/", handleSearchMove(apiKey))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func handleSearchMove(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")
		res, err := omdb.Search(apiKey, search)

		if err != nil {
			sendJSON(w, Response{Error: "something wrong with omdb"}, http.StatusBadGateway)
			return
		}

		sendJSON(w, Response{Data: res}, http.StatusOK)

	}
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal response", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)

	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response", "error", err)
	}
}
