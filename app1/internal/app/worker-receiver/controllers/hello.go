package controllers

import (
	"app1/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (lc *HelloController) Path() string {
	return "/hello"
}

func (rc *HelloController) Validate(ctx context.Context, r *http.Request) *utils.InvalidHTTPRequestError {
	return nil
}

func (lc *HelloController) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	responseJSON := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "OK",
		Message: "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(responseJSON); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
