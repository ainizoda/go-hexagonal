package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	inHttp "github.com/ainizoda/go-hexagonal/internal/adapters/in/http"
	"github.com/ainizoda/go-hexagonal/internal/adapters/in/http/handlers"
	"github.com/ainizoda/go-hexagonal/internal/adapters/out/memory"
	"github.com/ainizoda/go-hexagonal/internal/config"
	"github.com/ainizoda/go-hexagonal/internal/domain/user"
	"github.com/ainizoda/go-hexagonal/pkg/logger"
)

func main() {
	envPath := flag.String("env", ".env", "env path")
	flag.Parse()
	cfg, err := config.ParseConfig(*envPath)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	ur := memory.NewUserRepo()
	us := user.NewService(ur)
	uh := handlers.NewUserHandler(us)

	routes := []inHttp.Route{uh}
	server := inHttp.NewServer(cfg.Port, routes, lg)

	lg := logger.New(cfg.Env)
	lg.Info(context.Background(), fmt.Sprintf("server started at localhost:%d", cfg.Port))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Stop(shutdownCtx); err != nil {
			log.Fatalf("error shutting down server: %v", err)
		}
	}()
	if err := server.Start(); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
