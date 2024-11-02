package perrors

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	ctx := context.Background()
	err := errors.New("base error")

	err = Wrap(ctx, err, "augmented reason", nil)

	assert.Equal(t, "augmented reason: base error", err.Error())
}

func TestNestedCodes(t *testing.T) {
	ctx := context.Background()
	err := NewNotFound(ctx, "test not found", nil)

	err = Wrap(ctx, err, "internal error", nil)

	assert.Equal(t, "internal error: test not found", err.Error())
}
