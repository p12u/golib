package perrors

import (
	"context"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/fmsg"
	"github.com/Southclaws/fault/ftag"
)

func New(ctx context.Context, code ftag.Kind, message ...string) error {
	description := ""
	externalDescription := ""

	if len(message) > 1 {
		description = message[0]
	}

	if len(message) > 2 {
		externalDescription = message[1]
	}

	return fault.New(
		description,
		fctx.With(ctx),
		ftag.With(code),
		fmsg.WithDesc(description, externalDescription),
	)
}

func NewInternal(ctx context.Context, message ...string) error {
	return New(ctx, ftag.Internal, message...)
}

func NewNotFound(ctx context.Context, message ...string) error {
	return New(ctx, ftag.NotFound, message...)
}
