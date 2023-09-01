package services

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/tofu345/Building-mgmt-backend/src"
)

var v *validator.Validate

var validators = []CustomValidator{
	{
		"pswd", func(fl validator.FieldLevel) bool {
			return len(fl.Field().String()) >= 6 && len(fl.Field().String()) <= 30
		},
	},
}

type CustomValidator struct {
	name string
	f    func(validator.FieldLevel) bool
}

func init() {
	v = validator.New()

	for _, obj := range validators {
		err := v.RegisterValidation(obj.name, obj.f)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ValidateModel(obj any) error {
	return v.Struct(obj)
}

func ValidateMap(data map[string]any, schema map[string]any) map[string]any {
	return v.ValidateMap(data, schema)
}

func FmtValidationErrors(err error) map[string]string {
	switch err.(type) {
	case *validator.InvalidValidationError:
		fmt.Printf("! validator.InvalidValidationError")
		return map[string]string{"message": err.Error()}
	}

	errs := map[string]string{}
	var errMsg string

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errMsg = src.RequiredField
		case "email":
			errMsg = "Invalid email address"
		case "pswd":
			errMsg = "Password must be at 6-30 characters"
		default:
			errMsg = err.Tag()
		}

		errs[err.Field()] = errMsg
	}

	return errs
}
