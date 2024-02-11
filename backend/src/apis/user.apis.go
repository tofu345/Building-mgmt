package apis

import (
	"net/http"

	"github.com/tofu345/Building-mgmt-backend/src/constants"
	m "github.com/tofu345/Building-mgmt-backend/src/models"
	s "github.com/tofu345/Building-mgmt-backend/src/services"
)

func GenerateTokenPair(w http.ResponseWriter, r *http.Request) {
	data := s.GetContextData(r, "validated_data").(map[string]string)

	user, err := m.GetUserByEmail(data["email"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	if !s.CheckPasswordHash(data["password"], user.Password) {
		s.BadRequest(w, constants.InvalidLogin)
		return
	}

	access, err := s.AccessToken(user)
	if err != nil {
		s.BadRequest(w, constants.TokenError)
	}

	refresh, err := s.RefreshToken(user)
	if err != nil {
		s.BadRequest(w, constants.TokenError)
	}

	s.Success(w, map[string]any{"access": access, "refresh": refresh})
}

func RegenerateAccessToken(w http.ResponseWriter, r *http.Request) {
	data := s.GetContextData(r, "validated_data").(map[string]string)

	payload, err := s.DecodeToken(data["token"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	if _, exists := payload["ref"]; !exists {
		s.BadRequest(w, constants.InvalidToken)
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
			s.BadRequest(w, constants.TokenError)
			return
		}

		s.Success(w, map[string]any{"access": access})
	default:
		s.BadRequest(w, constants.InvalidToken)
	}
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	user := s.GetContextData(r, "user").(m.User)
	if !user.IsSuperuser {
		s.JsonResponse(w, http.StatusUnauthorized, constants.ErrUnauthorized)
		return
	}

	users, err := m.GetUserList()
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, users)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	user := s.GetContextData(r, "user").(m.User)
	s.Success(w, user)
}
