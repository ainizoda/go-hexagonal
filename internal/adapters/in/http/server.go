package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ainizoda/go-hexagonal/internal/adapters/in/http/middleware"
	"github.com/ainizoda/go-hexagonal/pkg/logger"
)

type Server struct {
	srv *http.Server
}

func NewServer(port int, routes []Route, logger *logger.L) *Server {
	mux := http.NewServeMux()

	for _, r := range routes {
		r.Register(mux)
	}

	handler := middleware.LoggingMiddleware(logger)(mux)
	handler = middleware.ErrorHandlingMiddleware(logger)(handler)

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      handler,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}
	return &Server{srv: srv}
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
