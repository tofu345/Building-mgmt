package internal

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Route struct {
	url      string
	methods  []string
	function func(http.ResponseWriter, *http.Request)
}

var Routes = []Route{
	{"/locations", []string{"GET"}, locations},
	{"/locations", []string{"POST"}, createLocation},
	{"/locations/{id}", []string{"PUT"}, updateLocation},
	{"/token", []string{"POST"}, getTokens},
	{"/token/refresh", []string{"POST"}, refreshToken},
}

func RegisterRoutes(r *mux.Router) {
	for _, route := range Routes {
		r.HandleFunc(route.url, route.function).Methods(route.methods...)
	}
}

func parseError(err error) string {
	str := err.Error()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		str = "Object not found"
	case strings.HasPrefix(str, "UNIQUE constraint failed: "):
		str = strings.Split(str, ": ")[1] + " is already in use"
	}

	return str
}

func jsonResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	out := map[string]any{}
	if code == http.StatusOK {
		out["responseCode"] = 100
	} else {
		out["responseCode"] = 103
	}

	switch data := data.(type) {
	case map[string]any:
		for k, v := range data {
			out[k] = v
		}
	default:
		out["data"] = data
	}

	json.NewEncoder(w).Encode(out)
}

func jsonError(w http.ResponseWriter, err any) {
	switch err := err.(type) {
	case error:
		jsonResponse(w, 400, map[string]any{"message": "An Error Occured", "error": parseError(err)})
		return
	}
	jsonResponse(w, 400, map[string]any{"message": "An Error Occured", "error": err})
}

func jsonDecode(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if errors.Is(err, io.EOF) {
		return ErrEmptyPostData
	}
	return err
}
