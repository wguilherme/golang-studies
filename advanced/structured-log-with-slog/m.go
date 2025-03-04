package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

// isso resulta em DEBUG-46, ou seja, 46 níveis abaixo do log level de DEBUG,
const FooLevel = slog.Level(-50)

func main() {

	opts := &slog.HandlerOptions{
		AddSource: true,
		// Level:     slog.LevelDebug,
		Level: FooLevel, // log level customizado
		// ReplaceAttr: nil, este campo daria para tratar o level do log, por ex.: um level de log customizado.
		// exemplo:
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "level" {
				level := a.Value.String()
				if level == "DEBUG-46" {
					a.Value = slog.StringValue("FOO")
				}
			}
			return a
		},
	}
	l := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(l)
	slog.Debug("foo")
	slog.Info("Service starting", "version", "1.0.0")
	l = l.With(slog.Group("app_info", slog.String("version", "1.0.0")))
	l.LogAttrs(context.Background(), FooLevel, "any message")
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
