package main

import (
	"log"

	inHttp "github.com/ainizoda/go-hexagonal/internal/adapters/in/http"
	"github.com/ainizoda/go-hexagonal/internal/adapters/in/http/handlers"
	"github.com/ainizoda/go-hexagonal/internal/adapters/in/http/server"
	"github.com/ainizoda/go-hexagonal/internal/adapters/out/memory"
	"github.com/ainizoda/go-hexagonal/internal/domain/user"
)

func main() {
	logger := log.Default()

	ur := memory.NewUserRepo()
	us := user.NewService(ur)
	uh := handlers.NewUserHandler(us)

	routes := []inHttp.Route{uh}

	server := server.NewServer(8080, routes, logger)

	logger.Println("Server successfully started at localhost:8080")

	if err := server.Start(); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
