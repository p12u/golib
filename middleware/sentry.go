package middleware

import (
	sentryhttp "github.com/getsentry/sentry-go/http"
)

var SentryMiddleware = sentryhttp.New(sentryhttp.Options{
	Repanic: true,
}).Handle
