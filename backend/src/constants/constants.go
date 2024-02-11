package constants

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	JWT_KEY = os.Getenv("JWT_KEY")
	ALLOWED_HOSTS = strings.Split(os.Getenv("ALLOWED_HOSTS"), ",")
}

var (
	JWT_KEY       string
	ALLOWED_HOSTS []string

	ErrInvalidToken  = errors.New(InvalidToken)
	ErrInvalidData   = errors.New(InvalidData)
	ErrEmptyPostData = errors.New("empty post data")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrUserNotFound  = errors.New("user not found")
)

const (
	JWT_ISSUER = "Building Management"

	Success = "Success"

	RequiredField = "this field is required"
	InvalidData   = "invalid data"
	InvalidLogin  = "incorrect email or password"
	ErrorMessage  = "an error occured"
	TokenError    = "error generating token"
	InvalidToken  = "invalid or missing token"
)
