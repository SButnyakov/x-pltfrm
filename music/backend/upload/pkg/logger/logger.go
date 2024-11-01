package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	sl *slog.Logger
}

func New() *Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return &Logger{sl: logger}
}

func (l *Logger) Info(msg string, args ...any) {
	l.sl.Info(msg, args...)
}

func (l *Logger) Error(msg string, err error) {
	l.sl.Error(msg, slog.Any("error", err))
}

func (l *Logger) Fatal(msg string, err error) {
	l.Error(msg, err)
	os.Exit(-1)
}
