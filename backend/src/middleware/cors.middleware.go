package middleware

import (
	"net/http"

	"github.com/tofu345/Building-mgmt-backend/src/constants"
)

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range constants.ALLOWED_HOSTS {
			w.Header().Set("Access-Control-Allow-Origin", v)
		}
		next.ServeHTTP(w, r)
	})
}
