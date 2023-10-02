package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Parser struct{}

func (Parser) Parse(err error) []string {
	var errors []string
	for _, fieldErr := range err.(validator.ValidationErrors) {

		var errMsg string
		switch fieldErr.Tag() {
		case "required":
			errMsg = fmt.Sprintf("%s is required", fieldErr.Field())
		case "required_without":
			errMsg = fmt.Sprintf("One of %s or %s is required", fieldErr.Field(), fieldErr.Param())
		case "min":
			errMsg = fmt.Sprintf("%s must be at least %s characters", fieldErr.Field(), fieldErr.Param())
		case "max":
			errMsg = fmt.Sprintf("The maximum characters for %s is %s", fieldErr.Field(), fieldErr.Param())
		case "email":
			errMsg = "Invalid Email"
		case "url":
			errMsg = fmt.Sprintf("Invalid url for %s", fieldErr.Field())
		case "oneof":
			errMsg = fmt.Sprintf("The field %s must be one of %s", fieldErr.Field(), fieldErr.Param())
		default:
			errMsg = fmt.Sprintf("%s failed validation: %s", fieldErr.Field(), fieldErr.Tag())
		}

		errors = append(errors, errMsg)
	}

	return errors
}
