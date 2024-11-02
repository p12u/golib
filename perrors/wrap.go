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

func Wrap(ctx context.Context, err error, message string, metadata map[string]any) error {
	code := ftag.Get(err)
	if code == "" {
		code = CodeInternal
	}

	return WrapWithCode(ctx, err, code, message, metadata)
}

func WrapWithCode(
	ctx context.Context,
	err error,
	code ftag.Kind,
	message string,
	metadata map[string]any,
) error {
	return WrapForExternal(ctx, err, code, message, "", metadata)
}

func WrapForExternal(
	ctx context.Context,
	err error,
	code ftag.Kind,
	message string,
	externalDescription string,
	metadata map[string]any,
) error {
	wrappers := []fault.Wrapper{}

	if ctx != nil {
		if metadata != nil {
			ctx = fctx.WithMeta(ctx, metadataToKv(metadata)...)
		}
		wrappers = append(wrappers, fctx.With(ctx))
	}

	if code != "" {
		wrappers = append(wrappers, ftag.With(code))
	}

	if message != "" {
		wrappers = append(wrappers, fmsg.WithDesc(message, externalDescription))
	}

	return fault.Wrap(
		err,
		wrappers...,
	)
}
