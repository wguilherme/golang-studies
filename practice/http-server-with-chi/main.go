package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	Username string
	/*
		int8 = 1 byte
		int16 = 2 bytes
		int32 = 4 bytes
		int64 = 8 bytes
	*/
	ID       int64 `json:"id,string"` // json:"id,string" -> converte o valor para string
	Role     string
	Password string `json:"-"` // "-" -> não serializa o campo (não envia o campo na resposta)
}

func main() {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	db := map[int64]User{
		1: {
			ID:       1,
			Username: "admin",
			Role:     "admin",
			Password: "admin",
		},
	}

	r.Group(func(r chi.Router) {
		r.Use(jsonMiddleware)
		r.Get("/users/{id}", handleGetUsers(db))
		// r.Post("/users", handlePostUsers(db))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(idStr, 10, 64)

		user, ok := db[id]
		if ok {
			data, err := json.Marshal(user)
			if err != nil {
				panic(err)
			}
			w.Write(data)
		}
	}
}
