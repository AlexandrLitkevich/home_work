package logger

import (
	"github.com/AlexandrLitkevich/home_work/hw12_13_14_15_calendar/internal/logger/handlers/slogpretty"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

const (
	local = "local"
	dev   = "dev"
	prod  = "prod"
)

func New() *Logger {
	levelLog := viper.Get("logger.level")
	slog.Warn("this log level", slog.Any("LOG LEVEL", levelLog))

	var log *slog.Logger

	switch levelLog {
	case local:
		log = setupPrettySlog()
	case dev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case prod:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return &Logger{
		log,
	}
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
