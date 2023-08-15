package internal

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Model interface {
	requiredFields() []string
}

func validate(obj Model) (map[string]any, bool) {
	var structMap map[string]any
	data, _ := json.Marshal(obj)
	json.Unmarshal(data, &structMap)

	errorsMap := map[string]any{}
	for _, field := range obj.requiredFields() {
		v, ok := structMap[field]
		if !ok {
			keys := make([]string, 0, len(structMap))
			for k := range structMap {
				keys = append(keys, k)
			}
			panic(fmt.Sprintf(
				"field '%s' does not exist on model fields are %v",
				field, strings.Join(keys, ", "),
			))
		}

		if !validateField(v) {
			errorsMap[field] = RequiredField
		}
	}

	return errorsMap, len(errorsMap) == 0
}

// validateField checks if the provided field is not nil or empty
func validateField(field any) bool {
	switch field := field.(type) {
	case string:
		return field != ""
	case int, uint, float64:
		return field == 0
	}

	panic(fmt.Sprintf("type %t does not have any validator set", field))
}
