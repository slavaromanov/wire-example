package main

import (
	"context"
	"github.com/slavaromanov/wire-example/app"
	"log/slog"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := slog.Default()
	app, err := app.NewApp(ctx)
	if err != nil {
		logger.Error("failed to initialize app", err)
		os.Exit(1)
	}
	if err := app.Run(ctx); err != nil {
		logger.Error("failed to run app", err)
		os.Exit(1)
	}
	cancel()
}
