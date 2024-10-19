package perrors

import (
	"context"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/fmsg"
	"github.com/Southclaws/fault/ftag"
)

/**
* perrors is an opinionated lib exposing functionality for creating
* more context aware error handlings. It uses an undelrying library
* so most of the functionality wraps that libraries object.
 */

func Wrap(ctx context.Context, err error, message ...string) error {
	return WrapWithCode(ctx, err, ftag.Internal, message...)
}

func WrapWithCode(ctx context.Context, err error, code ftag.Kind, message ...string) error {
	description := ""
	externalDescription := ""

	if len(message) > 1 {
		description = message[0]
	}

	if len(message) > 2 {
		externalDescription = message[1]
	}

	return wrap(ctx, err, &code, description, externalDescription)
}

func wrap(
	ctx context.Context,
	err error,
	code *ftag.Kind,
	description string,
	externalDescription string,
) error {
	wrappers := []fault.Wrapper{}

	if ctx != nil {
		wrappers = append(wrappers, fctx.With(ctx))
	}

	if code != nil {
		wrappers = append(wrappers, ftag.With(*code))
	}

	if description != "" {
		wrappers = append(wrappers, fmsg.WithDesc(description, externalDescription))
	}

	return fault.Wrap(
		err,
		wrappers...,
	)
}
