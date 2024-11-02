package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/p12u/golib/logger"
	"github.com/p12u/golib/perrors"
)

func ErrorLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log := logger.Extract(c.Request().Context())
			fields := map[string]any{
				"uri":    v.URI,
				"status": v.Status,
			}

			if v.Error != nil {
				externalMessge := perrors.GetExternal(v.Error)
				if externalMessge != "" {
					fields["external_message"] = externalMessge
				}

				fields["stackrace"] = perrors.Stacktrace(v.Error)

				metadata := perrors.GetMetadata(v.Error)
				fields["errMeta"] = metadata

				log.Error().
					Err(v.Error).
					Stack().
					Fields(fields).
					Msg("http request")
			} else {
				log.Info().
					Fields(fields).
					Msg("http request")
			}

			return nil
		},
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
	})
}
