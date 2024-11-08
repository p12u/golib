package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	z "github.com/Oudwins/zog"
	"github.com/p12u/golib/perrors"
	"github.com/samber/lo"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateRequestBody(body interface{}) error {
	return validator.New(validator.WithRequiredStructEnabled()).Struct(body)
}

type Parser interface {
	Parse(any, any, ...z.ParsingOption) z.ZogErrMap
}

type CanValidate interface {
	Validator() Parser
}

// Echo convenient handler functions

// Echo Handler with Body and Query validated
func EBQ[Body CanValidate, Query CanValidate](
	h func(c echo.Context, body *Body, query *Query) error,
) echo.HandlerFunc {
	// this assumes the body will contain json data
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return perrors.Wrap(nil, err, "failed reading request body", nil)
		}

		decodedBody := map[string]any{}
		if err := json.Unmarshal(body, &decodedBody); err != nil {
			return perrors.Wrap(ctx, err, "failed unmarshaling request body", nil)
		}

		var t Body
		errsMap := t.Validator().Parse(decodedBody, &t)
		if errsMap != nil {
			sanitizedErrs := z.Errors.SanitizeMap(errsMap)
			return echo.NewHTTPError(
				http.StatusBadRequest,
				errors.New("request body validation failed").Error(),
			).WithInternal(perrors.New(
				ctx,
				perrors.CodeValidationFailed,
				sanitizedErrs["$first"][0],
				map[string]any{"validationErrors": sanitizedErrs},
			))
		}

		// Parse Query
		var q Query
		queryParams := lo.MapEntries(
			c.QueryParams(),
			func(key string, value []string) (string, any) {
				if len(value) == 1 {
					return key, value[0]
				}

				return key, value
			},
		)

		errsMap = q.Validator().Parse(queryParams, &q)
		if errsMap != nil {
			sanitizedErrs := z.Errors.SanitizeMap(errsMap)
			return echo.NewHTTPError(
				http.StatusBadRequest,
				errors.New("query validation failed").Error(),
			).WithInternal(perrors.New(
				ctx,
				perrors.CodeValidationFailed,
				sanitizedErrs["$first"][0],
				map[string]any{"validationErrors": sanitizedErrs},
			))
		}

		return h(c, &t, &q)
	}
}

// Echo Handler with Query validated
func EQ[Query CanValidate](
	h func(c echo.Context, query *Query) error,
) echo.HandlerFunc {
	// this assumes the body will contain json data
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		// Parse Query
		var q Query
		queryParams := lo.MapEntries(
			c.QueryParams(),
			func(key string, value []string) (string, any) {
				if len(value) == 1 {
					return key, value[0]
				}

				return key, value
			},
		)

		errsMap := q.Validator().Parse(queryParams, &q)
		if errsMap != nil {
			sanitizedErrs := z.Errors.SanitizeMap(errsMap)
			return echo.NewHTTPError(
				http.StatusBadRequest,
				errors.New("query validation failed").Error(),
			).WithInternal(perrors.New(
				ctx,
				perrors.CodeValidationFailed,
				sanitizedErrs["$first"][0],
				map[string]any{"validationErrors": sanitizedErrs},
			))
		}

		return h(c, &q)
	}
}

// Echo Handler with Body validated
func EB[Body CanValidate](
	h func(c echo.Context, body *Body) error,
) echo.HandlerFunc {
	// this assumes the body will contain json data
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return perrors.Wrap(nil, err, "failed reading request body", nil)
		}

		decodedBody := map[string]any{}
		if err := json.Unmarshal(body, &decodedBody); err != nil {
			return perrors.Wrap(ctx, err, "failed unmarshaling request body", nil)
		}

		var t Body
		errsMap := t.Validator().Parse(decodedBody, &t)
		if errsMap != nil {
			sanitizedErrs := z.Errors.SanitizeMap(errsMap)
			return echo.NewHTTPError(
				http.StatusBadRequest,
				errors.New("request body validation failed").Error(),
			).WithInternal(perrors.New(
				ctx,
				perrors.CodeValidationFailed,
				sanitizedErrs["$first"][0],
				map[string]any{"validationErrors": sanitizedErrs},
			))
		}

		return h(c, &t)
	}
}
