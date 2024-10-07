package controllers

import (
	"app1/pkg/utils"
	"context"
	"net/http"
)

type ReadinessController struct {
}

func NewReadinessController() *ReadinessController {
	return &ReadinessController{}
}

func (rc *ReadinessController) Path() string {
	return "/_/health/readiness"
}

func (rc *ReadinessController) Validate(ctx context.Context, r *http.Request) *utils.InvalidHTTPRequestError {
	return nil
}

func (rc *ReadinessController) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
