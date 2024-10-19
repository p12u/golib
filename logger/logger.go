package logger

import (
	"context"
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Default(env string) *zerolog.Logger {
	defaultLogger := log.With().Caller().Logger()

	if env == "" {
		env = os.Getenv("GO_ENV")
	}
	if env == "" {
		env = "development"
	}

	if env == "development" {
		defaultLogger = defaultLogger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else if flag.Lookup("test.v") != nil {
		defaultLogger = defaultLogger.Level(zerolog.ErrorLevel)
	}

	return &defaultLogger
}

func Extract(ctx context.Context) *zerolog.Logger {
	logger := zerolog.Ctx(ctx)
	if logger.GetLevel() == zerolog.Disabled {
		return Default("")
	}

	return logger
}

func Wrap(ctx context.Context, logger *zerolog.Logger) context.Context {
	return logger.WithContext(ctx)
}
