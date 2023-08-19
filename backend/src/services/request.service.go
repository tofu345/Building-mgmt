package services

import (
	"context"
	"net/http"

	c "github.com/tofu345/Building-mgmt-backend/src/constants"
)

func AddDataToRequestContext(r *http.Request, key string, data any) {
	newContext := context.WithValue(r.Context(), key, data)
	newRequest := r.WithContext(newContext)
	*r = *newRequest
}

// Parses post data to map and checks if required fields are present (if any)
func PostDataToMap(r *http.Request, requiredFields ...string) (map[string]string, bool) {
	data := map[string]string{}
	err := JsonDecode(r, &data)
	if err != nil {
		return map[string]string{"error": err.Error()}, false
	}

	if len(requiredFields) == 0 {
		return data, true
	}

	errorsMap := map[string]string{}
	for _, v := range requiredFields {
		if _, exists := data[v]; !exists {
			errorsMap[v] = c.RequiredField
		}
	}

	if len(errorsMap) == 0 {
		return data, true
	} else {
		return errorsMap, false
	}
}
