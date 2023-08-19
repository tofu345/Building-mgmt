package middleware

import (
	"net/http"

	s "github.com/tofu345/Building-mgmt-backend/src/services"
)

func AuthRequired(w http.ResponseWriter, r *http.Request) error {
	token := r.Header.Get("Authorization")
	user, err := s.JwtAuth(token)
	if err != nil {
		return err
	}

	s.AddDataToRequestContext(r, "user", user)
	return nil
}
