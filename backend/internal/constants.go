package internal

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

const (
	RequiredField = "This field is required"

	ISSUER = "Building mgmt"

	InvalidData  = "Invalid data"
	InvalidLogin = "Incorrect email or password"

	ErrorMessage = "An Error Occured"

	TokenError   = "Error Generating Token"
	InvalidToken = "Invalid or missing Token"
)

var (
	JWT_KEY       string
	ALLOWED_HOSTS []string

	ErrInvalidToken  = errors.New(InvalidToken)
	ErrInvalidData   = errors.New(InvalidData)
	ErrEmptyPostData = errors.New("Empty post data")
)
