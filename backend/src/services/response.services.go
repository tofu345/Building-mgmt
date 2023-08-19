package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/tofu345/Building-mgmt-backend/src/constants"
	"gorm.io/gorm"
)

func JsonResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func Success(w http.ResponseWriter, data any) {
	JsonResponse(w, 200, data)
}

func ParseError(err error) string {
	var str string
	switch err := err.(type) {
	case *pgconn.PgError:
		str = err.Detail
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

func BadRequest(w http.ResponseWriter, err any) {
	data := map[string]any{"message": "An error occured"}

	switch err := err.(type) {
	case error:
		data["detail"] = ParseError(err)
	case map[string]any:
		for k, v := range err {
			data[k] = v
		}
	default:
		data["detail"] = err
	}

	JsonResponse(w, 400, data)
}

func JsonDecode(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if errors.Is(err, io.EOF) {
		return constants.ErrEmptyPostData
	}
	return err
}
