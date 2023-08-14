package internal

import (
	"errors"
	"os"
)

const (
	RequiredField = "This field is required"

	ISSUER = "Building mgmt"

	InvalidData  = "Invalid data"
	InvalidLogin = "Incorrect email or password"

	TokenError   = "Error Generating Token"
	InvalidToken = "Invalid or missing Token"
)

var (
	JWT_KEY = os.Getenv("JWT_KEY")

	ErrInvalidToken  = errors.New(InvalidToken)
	ErrInvalidData   = errors.New(InvalidData)
	ErrEmptyPostData = errors.New("Empty post data")
)
