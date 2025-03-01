package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type ApplicationHelper interface {
	RegisterGracefulShutdown(ctx context.Context, cancel context.CancelFunc)
}

func NewApplicationHelper() ApplicationHelper {
	return &applicationHelper{}
}

type applicationHelper struct {
}

func (ah applicationHelper) RegisterGracefulShutdown(ctx context.Context, cancel context.CancelFunc) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		select {
		case <-ctx.Done():
			GetLogger().Info(ctx).Msg("context ended")
		case sig := <-sigs:
			GetLogger().Info(ctx).
				Str("signal", sig.String()).
				Msg("received OS signal")
			cancel()
		}
	}()

	<-ctx.Done()

	GetLogger().Info(ctx).Msg("shutting down gracefully")

	GetLogger().Info(ctx).Msg("shutdown complete")
}
