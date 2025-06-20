package logger

import (
	"log/slog"
	"os"
)

func MustLoadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
}
