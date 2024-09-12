package safe

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	result := Ok(1)
	assert.True(t, result.IsOk())
	assert.False(t, result.IsErr())
	assert.Equal(t, 1, result.Unwrap())
	assert.Panics(t, func() {
		_ = result.UnwrapErr()
	})
}

func TestErr(t *testing.T) {
	result := Err[int](fmt.Errorf("error"))
	assert.False(t, result.IsOk())
	assert.True(t, result.IsErr())
	assert.Equal(t, "error", result.UnwrapErr().Error())
	assert.Panics(t, func() {
		result.Unwrap()
	})
}

func TestErr_Nil(t *testing.T) {
	result := Err[int](nil)
	assert.False(t, result.IsOk())
	assert.True(t, result.IsErr())
	assert.Equal(t, "Err called with nil error", result.UnwrapErr().Error())
}

func TestUnwrapFmt(t *testing.T) {
	result := Err[int](fmt.Errorf("error"))
	assert.Equal(t, "test: error", result.UnwrapFmt("test: %w").Error())
}

func TestUnwrapFmt_OK(t *testing.T) {
	result := Ok(1)
	assert.Panics(t, func() {
		_ = result.UnwrapFmt("test: %w")
	})
}
