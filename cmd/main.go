package main

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ekediala/expensix/config"
	"github.com/ekediala/expensix/httpio"
	"github.com/ekediala/expensix/server"
	"github.com/ekediala/expensix/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	const appName = "expensix"

	logger := slog.New(httpio.NewLogHandler(slog.NewJSONHandler(os.Stderr, nil)))
	logger = logger.With("app", appName)
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		logger.Error("main", "loading env variables", err)
		os.Exit(1)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("main", "loading config", err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	db, err := sqlx.Dial(ctx, cfg.DB.URL)
	if err != nil {
		logger.Error("main", "dialing db", err)
		return // we want the deferred function "stop" to run and free up the signals
	}

	store := sqlx.New(db)

	app := server.New(store)
	handler := http.TimeoutHandler(app, 5*time.Second, "request timeout")
	handler = httpio.CORSMiddleware(handler)
	handler = httpio.LoggingMiddleware(handler)
	handler = httpio.TraceMiddleware(handler)

	address := fmt.Sprintf("localhost:%s", cfg.Server.Port)
	server := http.Server{
		Handler: handler,
		Addr:    address,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.InfoContext(ctx, "main", "listening and serving", err)
		}
	}()

	logger.InfoContext(ctx, "main", "server running on", address)

	<-ctx.Done()
	ctx, done := context.WithTimeout(context.Background(), time.Second*30)
	defer done()

	err = server.Shutdown(ctx)
	logger.InfoContext(ctx, "main", "shutting down", cmp.Or(err, errors.New("server shutdown gracefully")))
}
