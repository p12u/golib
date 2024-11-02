package perrors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctx := context.Background()
	err := new(ctx, CodeInternal, "base error", nil)

	assert.Equal(t, "base error", err.Error())
}
