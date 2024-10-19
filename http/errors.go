package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const validatorErrorFmt = "field '%s' failed validation rule '%s'"

type Error struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type errorResponse struct {
	Errors []Error `json:"errors"`
}

// RequestErrorHandler is a function that handles errors that occur during request processing.
func RequestErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	encodeErr := json.NewEncoder(w).Encode(errorResponse{
		Errors: []Error{
			{
				Code:    "bad_request",
				Message: err.Error(),
			},
		},
	})

	if encodeErr != nil {
		http.Error(w, errors.Join(err, encodeErr).Error(), http.StatusInternalServerError)
	}
}

func ResponseErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	var statusCode int
	var errs []Error

	// Check for validation errors
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		statusCode = http.StatusBadRequest
		for _, fieldErr := range validationError {
			errs = append(errs, Error{
				Code:    "field_validation_error",
				Message: fmt.Sprintf(validatorErrorFmt, fieldErr.Field(), fieldErr.Tag()),
				Params: map[string]interface{}{
					"field":      fieldErr.Field(),
					"validation": fieldErr.Tag(),
				},
			})
		}
	} else {
		statusCode = http.StatusInternalServerError
		errs = []Error{
			{
				Code:    "internal_server_error",
				Message: err.Error(),
			},
		}
	}

	w.WriteHeader(statusCode)
	encodeErr := json.NewEncoder(w).Encode(errorResponse{
		Errors: errs,
	})
	if encodeErr != nil {
		http.Error(w, errors.Join(err, encodeErr).Error(), http.StatusInternalServerError)
	}
}
