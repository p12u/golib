package middleware

import (
	"context"
	"net/http"

	"github.com/p12u/golib/logger"
)

// OAPI strict handler
type StrictHandlerFunc func(context.Context, http.ResponseWriter, *http.Request, interface{}) (interface{}, error)

// Injects the logger into the request context
// It also includes request context into the logs
func OapiLoggerMiddleware(
	f StrictHandlerFunc,
	operationId string,
) StrictHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, args interface{}) (interface{}, error) {
		log := logger.Extract(ctx).With().Str("request_type", operationId).Logger()

		ctx = logger.Wrap(ctx, &log)
		result, err := f(ctx, w, r, args)
		if err != nil {
			log = log.With().Err(err).Logger()
		}

		log.Info().Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("request completed")

		return result, err
	}
}
