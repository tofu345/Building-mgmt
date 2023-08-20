package middleware

import (
	"net/http"

	s "github.com/tofu345/Building-mgmt-backend/src/services"
)

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		user, err := s.JwtAuth(token)
		if err != nil {
			s.JsonError(w, http.StatusUnauthorized, err)
			return
		}

		s.AddToRequestContext(r, "user", user)
		next.ServeHTTP(w, r)
	})
}
