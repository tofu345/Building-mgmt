package services

import (
	"log"

	"github.com/go-playground/validator"
	c "github.com/tofu345/Building-mgmt-backend/src/constants"
)

var v *validator.Validate

func init() {
	v = validator.New()
	err := v.RegisterValidation("pswd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 6 && len(fl.Field().String()) <= 30
	})
	if err != nil {
		log.Fatal(err)
	}
}

// func Validator() *validator.Validate {
// 	return v
// }

func ValidateModel(obj any) map[string]string {
	err := v.Struct(obj)
	if err == nil {
		return nil
	}

	errMap := map[string]string{}
	var errMsg string

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errMsg = c.RequiredField
		case "email":
			errMsg = "Invalid email address"
		case "pswd":
			errMsg = "Password must be at 6-30 characters"
		default:
			errMsg = err.Tag()
		}

		errMap[err.Field()] = errMsg
	}

	return errMap
}
