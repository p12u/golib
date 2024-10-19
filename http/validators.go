package http

import "github.com/go-playground/validator/v10"

func ValidateRequestBody(body interface{}) error {
	return validator.New(validator.WithRequiredStructEnabled()).Struct(body)
}
