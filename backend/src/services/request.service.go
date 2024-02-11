package services

import (
	"context"
	"net/http"
)

type Key string // To avoid collisions

func AddToRequestContext(r *http.Request, key string, data any) {
	newContext := context.WithValue(r.Context(), Key(key), data)
	newRequest := r.WithContext(newContext)
	*r = *newRequest
}

func GetContextData(r *http.Request, key string) any {
	data := r.Context().Value(Key(key))
	return data
}
