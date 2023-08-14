package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func jwtAuth(r *http.Request) (User, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return User{}, ErrInvalidToken
	}

	token = strings.Split(token, " ")[1]

	payload, err := decodeToken(token)
	if err != nil {
		return User{}, ErrInvalidToken
	}

	email := payload["email"]
	switch email := email.(type) {
	case string:
		var user User
		err := db.First(&user, "email = ?", email).Error
		if err != nil {
			return User{}, err
		}

		return user, nil
	default:
		return User{}, ErrInvalidToken
	}
}

func newAccessToken(user User) (string, error) {
	key := []byte(JWT_KEY)
	time_now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   ISSUER,
		"iat":   time_now.Unix(),
		"exp":   time_now.Add(time.Hour).Unix(),
		"email": user.Email,
	})
	return t.SignedString(key)
}

func newRefreshToken(user User) (string, error) {
	key := []byte(JWT_KEY)
	time_now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   ISSUER,
		"iat":   time_now.Unix(),
		"exp":   time_now.Add(time.Hour * 24).Unix(),
		"ref":   true,
		"email": user.Email,
	})
	return t.SignedString(key)
}

func decodeToken(tokenData string) (map[string]any, error) {
	token, err := jwt.Parse(tokenData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_KEY), nil
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
