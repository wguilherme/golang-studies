package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(l)
	slog.Debug("foo")
	slog.Info("Service starting", "version", "1.0.0")
	l = l.With(slog.Group("app_info", slog.String("version", "1.0.0")))
	l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"tivemos um http request",
		slog.Group("http_data", slog.String(
			"method",
			http.MethodDelete,
		),
			slog.Int("status", http.StatusOK),
			slog.String("user_agent", "hasuida"),
		))

}
