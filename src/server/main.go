package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// By default, set to text for local dev, for prod build make sure to use json.
var logFormatter = "text"

func main() {
	// test logger
	logger := getLogger(logFormatter)
	logger.Info("starting server")

	ctx := context.Background()

	if err := run(ctx, logger); err != nil {
		logger.Fatal("run function unexpected exit", "err", err)
	}
}

func run(ctx context.Context, logger *log.Logger) error {
	//router := chi.NewRouter()
	return fmt.Errorf("testing error")
}

// getLogger creates a new logger to be used based on the logFormatter variable
func getLogger(format string) *log.Logger {
	var formatter log.Formatter
	switch format {
	case "text":
		formatter = log.TextFormatter
	default:
		// use jsonformatter whenever the logFormatter variable is set differently
		formatter = log.JSONFormatter
	}
	logger := log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat:      time.RFC3339,
		Prefix:          "Samtala A.S.",
		ReportTimestamp: true,
		ReportCaller:    true,
		Formatter:       formatter,
	})
	return logger
}
