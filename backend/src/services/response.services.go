package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/tofu345/Building-mgmt-backend/src"
	"gorm.io/gorm"
)

func ParseError(err error) any {
	var str string
	switch err := err.(type) {
	case *pgconn.PgError:
		str = err.Detail
	case validator.ValidationErrors:
		return FmtValidationErrors(err)
	default:
		str = err.Error()

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			str = "Object not found"
		case strings.HasPrefix(str, "UNIQUE constraint failed: "):
			str = strings.Split(str, ": ")[1] + " is already in use"
		}
	}

	return str
}

func JsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Success(w http.ResponseWriter, data any) {
	switch data := data.(type) {
	case string:
		JsonResponse(w, http.StatusOK, map[string]string{"message": data})
	default:
		JsonResponse(w, http.StatusOK, data)
	}
}

func JsonError(w http.ResponseWriter, status int, err error) {
	data := map[string]any{
		"message": "An error occured",
		"error":   ParseError(err),
	}
	JsonResponse(w, status, data)
}

func BadRequest(w http.ResponseWriter, err any) {
	data := map[string]any{"message": "An error occured"}

	switch err := err.(type) {
	case error:
		data["detail"] = ParseError(err)
	case map[string]any:
		data["errors"] = err
	default:
		data["detail"] = err
	}

	JsonResponse(w, http.StatusBadRequest, data)
}

func JsonDecode(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if errors.Is(err, io.EOF) {
		return src.ErrEmptyPostData
	}
	return err
}
