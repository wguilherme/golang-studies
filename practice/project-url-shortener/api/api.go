package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewHandler(db map[string]string) http.Handler {
	// inicializa o router
	r := chi.NewMux()

	// aplica middlewares
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// declara as rotas
	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(
				w,
				Response{Error: "invalid body"},
				http.StatusUnprocessableEntity)
			return
		}
		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(
				w,
				Response{Error: "invalid url"},
				http.StatusBadRequest,
			)
			return
		}

		code := genCode()
		db[code] = body.URL
		sendJSON(w, Response{Data: code}, http.StatusCreated)

	}
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		code := chi.URLParam(r, "code")
		data, ok := db[code]
		if !ok {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, data, http.StatusPermanentRedirect)

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

const CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genCode() string {
	const n = 8
	byts := make([]byte, n)
	for i := range n {
		byts[i] = CHARACTERS[rand.IntN(len(CHARACTERS))]
	}
	return string(byts)
}
