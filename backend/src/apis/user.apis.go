package apis

import (
	"net/http"

	c "github.com/tofu345/Building-mgmt-backend/src/constants"
	m "github.com/tofu345/Building-mgmt-backend/src/models"
	s "github.com/tofu345/Building-mgmt-backend/src/services"
)

func GenerateTokenPair(w http.ResponseWriter, r *http.Request) {
	data, valid := s.PostDataToMap(r, "email", "password")
	if !valid {
		s.BadRequest(w, data)
		return
	}

	user, err := m.GetUserByEmail(data["email"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	if !s.CheckPasswordHash(data["password"], user.Password) {
		s.BadRequest(w, c.InvalidLogin)
		return
	}

	access, err := s.AccessToken(user)
	if err != nil {
		s.BadRequest(w, c.TokenError)
	}

	refresh, err := s.RefreshToken(user)
	if err != nil {
		s.BadRequest(w, c.TokenError)
	}

	s.Success(w, map[string]any{"access": access, "refresh": refresh})
}

func RegenerateAccessToken(w http.ResponseWriter, r *http.Request) {
	data, valid := s.PostDataToMap(r, "email", "password")
	if !valid {
		s.BadRequest(w, data)
		return
	}

	payload, err := s.DecodeToken(data["token"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	if _, exists := payload["ref"]; !exists {
		s.BadRequest(w, c.InvalidToken)
		return
	}

	email := payload["email"]
	switch email := email.(type) {
	case string:
		user, err := m.GetUserByEmail(email)
		if err != nil {
			s.BadRequest(w, err)
			return
		}

		access, err := s.AccessToken(user)
		if err != nil {
			s.BadRequest(w, c.TokenError)
			return
		}

		s.Success(w, map[string]any{"access": access})
	default:
		s.BadRequest(w, c.InvalidToken)
	}
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	users, err := m.GetUserList()
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, users)
}
