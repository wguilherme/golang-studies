package utils

import (
	"context"
	"net/http"
	"time"
)

type HTTPController interface {
	Path() string
	Validate(ctx context.Context, r *http.Request) *InvalidHTTPRequestError
	Handle(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type HTTPHelper interface {
	CreateHandler(ctx context.Context, controller HTTPController) http.HandlerFunc
	RegisterControllers(ctx context.Context, controllers ...HTTPController) []http.HandlerFunc
	SendRequest(ctx context.Context, req *http.Request, client *http.Client) (*http.Response, error)
}

func NewHTTPHelper() HTTPHelper {
	return &httpHelper{}
}

type httpHelper struct {
}

func (h *httpHelper) CreateHandler(ctx context.Context, controller HTTPController) http.HandlerFunc {

	GetLogger().Debug(ctx).
		Str("path", controller.Path()).
		Msg("creating http handler")

	handler := func(w http.ResponseWriter, r *http.Request) {

		var c = ContextWithTransaction(r.Context())

		var cancel context.CancelFunc

		timeout, err := time.ParseDuration(r.FormValue("timeout"))
		if err == nil {
			GetLogger().Debug(c).
				Str("timeout", timeout.String()).
				Msg("context with timeout")
			c, cancel = context.WithTimeout(c, timeout)
		} else {
			GetLogger().Debug(c).
				Msg("context with cancel")
			c, cancel = context.WithCancel(c)
		}

		defer cancel()

		GetLogger().Debug(c).
			Str("path", controller.Path()).
			Str("method", r.Method).
			Msg("new http request received")

		if e := controller.Validate(c, r); e != nil {
			http.Error(w, e.Error(), e.StatusCode)
			return
		}

		controller.Handle(c, w, r)

		GetLogger().Debug(c).
			Str("path", controller.Path()).
			Str("method", r.Method).
			Msg("new http request processed")

	}

	GetLogger().Debug(ctx).
		Str("path", controller.Path()).
		Msg("http handler created")

	return handler

}

func (h *httpHelper) RegisterControllers(ctx context.Context, controllers ...HTTPController) []http.HandlerFunc {

	handlers := []http.HandlerFunc{}

	for _, controller := range controllers {

		GetLogger().Debug(ctx).
			Str("path", controller.Path()).
			Msg("registering http controller")

		handler := h.CreateHandler(ctx, controller)

		http.HandleFunc(controller.Path(), handler)

		GetLogger().Debug(ctx).
			Str("path", controller.Path()).
			Msg("http controller was successfully registered")

		handlers = append(handlers, handler)

	}

	return handlers

}

func (h *httpHelper) SendRequest(ctx context.Context, req *http.Request, client *http.Client) (*http.Response, error) {

	GetLogger().Debug(ctx).
		Msg("sending http request")

	if client == nil {
		client = &http.Client{}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	GetLogger().Debug(ctx).
		Msg("http request was successfully sent")

	return resp, nil

}

type InvalidHTTPRequestError struct {
	StatusCode int
	Err        error
}

func NewInvalidHTTPRequestError(statusCode int, err error) *InvalidHTTPRequestError {
	return &InvalidHTTPRequestError{
		StatusCode: statusCode,
		Err:        err,
	}
}

func (e *InvalidHTTPRequestError) Error() string {
	return e.Err.Error()
}
