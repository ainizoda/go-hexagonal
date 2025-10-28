package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	inHttp "github.com/ainizoda/go-hexagonal/internal/adapters/in/http"
	"github.com/ainizoda/go-hexagonal/internal/adapters/in/http/handlers"
	"github.com/ainizoda/go-hexagonal/internal/adapters/out/memory"
	"github.com/ainizoda/go-hexagonal/internal/config"
	"github.com/ainizoda/go-hexagonal/internal/domain/user"
	"github.com/ainizoda/go-hexagonal/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	envPath := flag.String("env", ".env", "path to env file")
	flag.Parse()

	cfg, err := config.ParseConfig(*envPath)
	if err != nil {
		return fmt.Errorf("error reading config: %w", err)
	}

	lg := logger.New(cfg.Env)

	userRepo := memory.NewUserRepo()
	userService := user.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routes := []inHttp.Route{userHandler}
	server := inHttp.NewServer(cfg.Port, routes, lg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	lg.Info(ctx, "server starting", zap.Int("port", cfg.Port))

	// Start server in background
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Start()
	}()

	select {
	case <-ctx.Done():
		lg.Info(ctx, "shutdown signal received, stopping server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Stop(shutdownCtx); err != nil {
			return fmt.Errorf("error shutting down server: %w", err)
		}
		lg.Info(ctx, "server stopped gracefully")
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("server error: %w", err)
		}
	}

	return nil
}
