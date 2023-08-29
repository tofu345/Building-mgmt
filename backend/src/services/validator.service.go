package services

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/tofu345/Building-mgmt-backend/src"
)

var v *validator.Validate

var validators = []CustomValidator{
	{"pswd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 6 && len(fl.Field().String()) <= 30
	}},
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

func Validator() *validator.Validate {
	return v
}

func ValidateModel(obj any) map[string]string {
	err := v.Struct(obj)
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case *validator.InvalidValidationError:
		log.Println("! Invalid Validation Error", obj)
		return nil
	case validator.ValidationErrors:
		return FormatValidationErrors(err)
	default:
		log.Printf("! Unknown validation error of type %T", obj)
		return nil
	}
}

func FormatValidationErrors(err validator.ValidationErrors) map[string]string {
	errs := map[string]string{}
	var errMsg string

	for _, err := range err {
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
