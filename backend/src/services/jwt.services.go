package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tofu345/Building-mgmt-backend/src"
	m "github.com/tofu345/Building-mgmt-backend/src/models"
)

func JwtAuth(token string) (m.User, error) {
	if token == "" {
		return m.User{}, src.ErrInvalidToken
	}

	token = strings.Split(token, " ")[1]
	payload, err := DecodeToken(token)
	if err != nil {
		return m.User{}, src.ErrInvalidToken
	}

	email := payload["email"]
	switch email := email.(type) {
	case string:
		user, err := m.GetUserByEmail(email)
		if err != nil {
			return m.User{}, err
		}

		return user, nil
	default:
		return m.User{}, src.ErrInvalidToken
	}
}

func defaultJwtClaims(user m.User) jwt.MapClaims {
	time_now := time.Now()
	return jwt.MapClaims{
		"iss":   src.JWT_ISSUER,
		"iat":   time_now.Unix(),
		"exp":   time_now.Add(time.Hour).Unix(),
		"email": user.Email,
	}
}

func AccessToken(user m.User) (string, error) {
	key := []byte(src.JWT_KEY)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, defaultJwtClaims(user))
	return t.SignedString(key)
}

func RefreshToken(user m.User) (string, error) {
	key := []byte(src.JWT_KEY)
	claims := defaultJwtClaims(user)
	claims["ref"] = true
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(key)
}

func DecodeToken(tokenData string) (map[string]any, error) {
	token, err := jwt.Parse(tokenData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(src.JWT_KEY), nil
	})
	if err != nil {
		return map[string]any{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return map[string]any{}, err
	}
}
