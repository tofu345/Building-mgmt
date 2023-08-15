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
	url          string
	methods      []string
	function     func(http.ResponseWriter, *http.Request)
	authRequired bool
}

var Routes = []Route{
	{url: "/locations", methods: []string{"GET"}, function: locations, authRequired: true},
	{url: "/locations", methods: []string{"POST"}, function: createLocation, authRequired: true},
	{url: "/locations/{id}", methods: []string{"PUT"}, function: updateLocation, authRequired: true},

	{url: "/token", methods: []string{"POST"}, function: getTokens},
	{url: "/token/refresh", methods: []string{"POST"}, function: refreshToken},
}

func RegisterRoutes(r *mux.Router) {
	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.Use(authRequiredMiddleware)

	for _, route := range Routes {
		if route.authRequired {
			authRouter.HandleFunc(route.url, route.function).Methods(route.methods...)
		} else {
			r.HandleFunc(route.url, route.function).Methods(route.methods...)
		}
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
		jsonResponse(w, 400, map[string]any{"message": ErrorMessage, "error": parseError(err)})
		return
	}
	jsonResponse(w, 400, map[string]any{"message": ErrorMessage, "error": err})
}

func jsonDecode(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if errors.Is(err, io.EOF) {
		return ErrEmptyPostData
	}
	return err
}
