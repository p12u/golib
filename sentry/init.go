package sentry

import (
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/p12u/golib/perrors"
)

type Options sentry.ClientOptions

func MustInit(opts *Options) func(time.Duration) {
	if opts == nil {
		opts = &Options{}
	}

	if opts.Dsn == "" {
		dsn, exist := os.LookupEnv("SENTRY_DSN")
		if !exist || dsn == "" {
			return func(d time.Duration) {}
		}
		opts.Dsn = dsn
	}

	err := sentry.Init(sentry.ClientOptions(*opts))
	if err != nil {
		panic(perrors.NewInternal(nil, "failed to init sentry", nil))
	}

	return func(d time.Duration) {
		sentry.Flush(d)
	}
}
