package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ainizoda/go-hexagonal/pkg/logger"
	"github.com/google/uuid"
)

func LoggingMiddleware(lg *logger.L) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			lw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(lw, r)

			ctx := logger.WithRequestID(r.Context(), uuid.NewString())
			lg.Debug(ctx, fmt.Sprintf("%s %s %d %s", r.Method, r.URL.Path, lw.statusCode, time.Since(start)))
		})
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lw *loggingResponseWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.ResponseWriter.WriteHeader(code)
}
