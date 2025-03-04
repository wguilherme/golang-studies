package main

import (
	"log/slog"
)

type User struct {
	Username string
	ID       int64 `json:"id,string"`
	Role     string
	Password string `json:"password"`
}

type Password string

func (u User) LogValue() slog.Value {
	// return slog.Int64(u.ID) também poderia retornar um grupo, ex.:
	return slog.GroupValue(
		slog.Int64("id", u.ID),
		slog.String("role", u.Role),
	)
}

func main() {

	u := User{
		Username: "admin",
		ID:       1,
		Role:     "admin",
		Password: "123456",
	}

	slog.Info("user", "u", u)

}
