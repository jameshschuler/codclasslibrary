package common

import "github.com/go-playground/validator/v10"

func ParseValidationErrors(err error) map[string]string {
	validationErrors := make(map[string]string)

	for _, fieldError := range err.(validator.ValidationErrors) {
		validationErrors[fieldError.Field()] = fieldError.Tag()
	}

	return validationErrors
}
