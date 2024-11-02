package perrors

import (
	"context"
	"errors"
	"fmt"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/ftag"
)

func New(
	ctx context.Context,
	code ftag.Kind,
	message string,
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

	return fault.Wrap(errors.New(message), wrappers...)
}

func NewInternal(ctx context.Context, message string, metadata map[string]any) error {
	return New(ctx, ftag.Internal, message, metadata)
}

func NewNotFound(ctx context.Context, message string, metadata map[string]any) error {
	return New(ctx, ftag.NotFound, message, metadata)
}

func metadataToKv(metadata map[string]any) []string {
	keyvalues := make([]string, 0, len(metadata))
	for k, v := range metadata {
		keyvalues = append(keyvalues, k, fmt.Sprintf("%v", v))
	}

	return keyvalues
}
