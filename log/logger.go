package log

import (
	"log/slog"
	"os"
)

var (
	logLevel   = new(slog.LevelVar)
	logOptions = &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: true,
	}
)

func init() {
	handler := slog.NewTextHandler(os.Stderr, logOptions)
	slog.SetDefault(slog.New(handler))
	logLevel.Set(slog.LevelInfo)
}

func ErrExit(err error, rc int) {
	slog.Error(err.Error())
	os.Exit(rc)
}
