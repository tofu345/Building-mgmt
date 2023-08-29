package services

import (
	"context"
	"net/http"
)

func AddToRequestContext(r *http.Request, key string, data any) {
	newContext := context.WithValue(r.Context(), key, data)
	newRequest := r.WithContext(newContext)
	*r = *newRequest
}

func GetContextData(r *http.Request, key string) any {
	data := r.Context().Value(key)
	return data
}
