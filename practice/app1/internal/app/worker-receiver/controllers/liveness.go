package controllers

import (
	"app1/pkg/utils"
	"context"
	"net/http"
)

type LivenessController struct {
}

func NewLivenessController() *LivenessController {
	return &LivenessController{}
}

func (lc *LivenessController) Path() string {
	return "/_/health/liveness"
}

func (rc *LivenessController) Validate(ctx context.Context, r *http.Request) *utils.InvalidHTTPRequestError {
	return nil
}

func (lc *LivenessController) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
