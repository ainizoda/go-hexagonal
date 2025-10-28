package main

import (
	"context"
	"flag"
	"fmt"
	"log"

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

	lg := logger.New(cfg.Env)

	ur := memory.NewUserRepo()
	us := user.NewService(ur)
	uh := handlers.NewUserHandler(us)

	routes := []inHttp.Route{uh}
	server := inHttp.NewServer(cfg.Port, routes, lg)

	lg.Info(context.Background(), fmt.Sprintf("server started at localhost:%d", cfg.Port))

	if err := server.Start(); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
