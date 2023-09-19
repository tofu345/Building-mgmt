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

func ValidateMap(data map[string]any, schema map[string]any) map[string][]string {
	data = v.ValidateMap(data, schema)
	return FmtValidationErrorsMap(data)
}

func FmtValidationErrors(errs error) map[string]string {
	switch errs := errs.(type) {
	case *validator.InvalidValidationError:
		fmt.Printf("! validator.InvalidValidationError")
		return map[string]string{"message": errs.Error()}
	}

	validation_errs, ok := errs.(validator.ValidationErrors)
	if !ok {
		panic(fmt.Sprintf("%T is not of type validator.ValidationErrors\n", errs))
	}

	err_map := make(map[string]string, len(validation_errs))
	for _, err := range validation_errs {
		k, v := _fmtValidationError(err)
		err_map[k] = v
	}

	return err_map
}

func _fmtValidationError(err validator.FieldError) (key string, value string) {
	switch err.Tag() {
	case "required":
		value = src.RequiredField
	case "email":
		value = "Invalid email address"
	case "pswd":
		value = "Password must be at 6-30 characters"
	default:
		value = err.Tag()
	}

	key = err.Field()
	return key, value
}

func FmtValidationErrorsMap(errs map[string]any) map[string][]string {
	output := make(map[string][]string, len(errs))

	for key, err := range errs {
		validation_errs, ok := err.(validator.ValidationErrors)
		if !ok {
			panic(fmt.Sprintf("%T is not of type validator.ValidationErrors\n", err))
		}

		err_slice := []string{}
		for _, v := range validation_errs {
			_, v := _fmtValidationError(v)
			err_slice = append(err_slice, v)

		}

		output[key] = err_slice
	}

	return output
}
