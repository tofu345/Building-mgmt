package middleware

import (
	"fmt"
	"net/http"

	s "github.com/tofu345/Building-mgmt-backend/src/services"
)

type Handler func(http.ResponseWriter, *http.Request)

func ValidateSchema(schema map[string]any, next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{}
		s.JsonDecode(r, &data)

		errs := s.ValidateMap(data, schema)
		if len(errs) > 0 {
			s.BadRequest(w, errs)
			return
		}

		validated_data := map[string]string{}
		for k, v := range data {
			validated_data[k] = fmt.Sprint(v)
		}

		s.AddToRequestContext(r, "validated_data", validated_data)
		http.HandlerFunc(next).ServeHTTP(w, r)
	}
}
