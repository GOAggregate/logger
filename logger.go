package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/GOAggregate/logger/handlers/slogpretty"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
)

func Init(logLevel LogLevel, logFile string) *slog.Logger {
	var log *slog.Logger

	if logFile == "" {
		switch logLevel {
		case LogLevelDebug:
			slogpretty.NewPrettyHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
		case LogLevelInfo:
			slogpretty.NewPrettyHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
		default:
			panic("log level not supported. Supported levels: debug, info")
		}
	}

	switch logLevel {
	case LogLevelDebug:
		log = slog.New(
			slog.NewJSONHandler(getFileToWrite(logFile), &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case LogLevelInfo:
		log = slog.New(
			slog.NewJSONHandler(getFileToWrite(logFile), &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		panic("log level not supported. Supported levels: debug, info")
	}

	log.Info("Logger init success")

	return log
}

func getFileToWrite(path string) io.Writer {
	if path == "" {
		return os.Stdout
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("failed open log file: " + path + " error: " + err.Error())
	}

	return f
}
