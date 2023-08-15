package internal

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func authRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := jwtAuth(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		newContext := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(newContext)
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
