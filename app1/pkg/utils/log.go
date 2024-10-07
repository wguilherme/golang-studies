package utils

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var logger TransactionalLoggerHelper

func GetLogger() TransactionalLoggerHelper {
	if logger == nil {
		logger = NewTransactionalLoggerHelper()
	}
	return logger
}

type TransactionalLoggerHelper interface {
	Info(ctx context.Context) *zerolog.Event
	Warn(ctx context.Context) *zerolog.Event
	Debug(ctx context.Context) *zerolog.Event
	Trace(ctx context.Context) *zerolog.Event
	Fatal(ctx context.Context) *zerolog.Event
	Error(ctx context.Context, err error) *zerolog.Event
	ErrorMessage(ctx context.Context, message string) *zerolog.Event
}

func NewTransactionalLoggerHelper() TransactionalLoggerHelper {

	ctx := ContextWithTransaction(context.Background())

	loggingLevel := flag.String("logging_level", zerolog.InfoLevel.String(), "logging level [-logging_level=trace|debug|info|error] (default is info)")
	loggingFormat := flag.String("logging_format", "json", "logging format [-logging_format=standard|json] (default is json)")

	flag.Parse()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	result := &transactionalLoggerHelper{
		loggerLevel: map[string]zerolog.Level{
			zerolog.TraceLevel.String(): zerolog.TraceLevel,
			zerolog.DebugLevel.String(): zerolog.DebugLevel,
			zerolog.InfoLevel.String():  zerolog.InfoLevel,
			zerolog.ErrorLevel.String(): zerolog.ErrorLevel,
		},
	}

	result.loadLoggingFormat(ctx, *loggingFormat)
	result.loadLoggingLevel(ctx, *loggingLevel)

	return result

}

type transactionalLoggerHelper struct {
	logger      zerolog.Logger
	loggerLevel map[string]zerolog.Level
}

func (l *transactionalLoggerHelper) loadLoggingFormat(ctx context.Context, loggingFormat string) {

	switch loggingFormat {

	case "json":

		l.logger = zerolog.New(os.Stderr).With().Logger()

		l.Info(ctx).
			Str("format", "json").
			Msg("setting logging format")

	default:

		l.logger = zerolog.New(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
			},
		).With().Timestamp().Logger()

		l.Info(ctx).
			Str("format", "standard").
			Msg("setting logging format")

	}

}

func (l *transactionalLoggerHelper) loadLoggingLevel(ctx context.Context, loggingLevel string) {

	level, ok := l.loggerLevel[loggingLevel]

	if ok {
		zerolog.SetGlobalLevel(level)
	} else {
		l.Info(ctx).
			Str("loggingLevelName", loggingLevel).
			Msg("unknown logging level, using default")
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	l.Info(ctx).
		Str("loggingLevel", zerolog.GlobalLevel().String()).
		Msg("logging level configured")

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		l.Info(ctx).Msg("stack enabled for logging")
		l.logger = l.logger.With().Stack().Logger()
	}

	if zerolog.GlobalLevel() == zerolog.TraceLevel {
		l.Info(ctx).Msg("caller enabled for logging")
		l.logger = l.logger.With().Stack().Caller().Logger()
	}

}

func (l *transactionalLoggerHelper) Info(ctx context.Context) *zerolog.Event {
	return l.logger.
		Info().
		Str(string(ContextTransactionKey), GetTransaction(ctx))
}

func (l *transactionalLoggerHelper) Warn(ctx context.Context) *zerolog.Event {
	return l.logger.
		Warn().
		Str(string(ContextTransactionKey), GetTransaction(ctx))
}

func (l *transactionalLoggerHelper) Debug(ctx context.Context) *zerolog.Event {
	return l.logger.
		Debug().
		Str(string(ContextTransactionKey), GetTransaction(ctx))
}

func (l *transactionalLoggerHelper) Trace(ctx context.Context) *zerolog.Event {
	return l.logger.
		Trace().
		Str(string(ContextTransactionKey), GetTransaction(ctx))
}

func (l *transactionalLoggerHelper) Fatal(ctx context.Context) *zerolog.Event {
	return l.logger.
		Fatal().
		Str(string(ContextTransactionKey), GetTransaction(ctx))
}

func (l *transactionalLoggerHelper) Error(ctx context.Context, err error) *zerolog.Event {
	wrappedError := errors.Wrap(err, err.Error())
	return l.logger.
		Error().
		Str(string(ContextTransactionKey), GetTransaction(ctx)).
		Err(wrappedError)
}

func (l *transactionalLoggerHelper) ErrorMessage(ctx context.Context, message string) *zerolog.Event {
	wrappedError := errors.New(message)
	return l.logger.
		Error().
		Str(string(ContextTransactionKey), GetTransaction(ctx)).Err(wrappedError)
}
