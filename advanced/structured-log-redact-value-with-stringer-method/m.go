package main

import (
	"log/slog"
)

type User struct {
	Username string
	ID       int64 `json:"id,string"`
	Role     string
	Password Password `json:"password"`
}

type Password string

func (p Password) String() string {
	return "[REDACTED]"
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
