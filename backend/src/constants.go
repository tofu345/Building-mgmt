package src

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

	ErrInvalidToken   = errors.New(InvalidToken)
	ErrInvalidData    = errors.New(InvalidData)
	ErrEmptyPostData  = errors.New("Empty post data")
	ErrObjectNotFound = errors.New("Object not found")
	ErrUnauthorized   = errors.New("Unauthorized")
)

const (
	JWT_ISSUER = "Building Management"

	RequiredField = "This field is required"
	InvalidData   = "Invalid data"
	InvalidLogin  = "Incorrect email or password"
	ErrorMessage  = "An error occured"
	TokenError    = "Error generating token"
	InvalidToken  = "Invalid or missing token"
)
