package middleware

import (
	"fmt"
	"net/http"

	"github.com/ainizoda/go-hexagonal/pkg/logger"
)

func ErrorHandlingMiddleware(logger *logger.L) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(r.Context(), fmt.Sprintf("PANIC: %v", err))
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
