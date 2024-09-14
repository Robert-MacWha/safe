package safe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapOption(t *testing.T) {
	result := MapOption(Some(1), func(i int) int {
		return i + 1
	})
	assert.True(t, result.IsSome())
	assert.Equal(t, 2, result.Unwrap())
}

func TestMapOption_None(t *testing.T) {
	result := MapOption(None[int](), func(i int) int {
		return i + 1
	})
	assert.True(t, result.IsNone())
}

func TestMapOptionOr(t *testing.T) {
	result := MapOptionOr(Some(1), 0, func(i int) int {
		return i + 1
	})
	assert.Equal(t, 2, result)
}

func TestMapOptionOr_None(t *testing.T) {
	result := MapOptionOr(None[int](), 0, func(i int) int {
		return i + 1
	})
	assert.Equal(t, 0, result)
}

func TestMapResult(t *testing.T) {
	result := MapResult(Ok(1), func(i int) int {
		return i + 1
	})
	assert.True(t, result.IsOk())
	assert.Equal(t, 2, result.Unwrap())
}

func TestMapResult_Err(t *testing.T) {
	result := MapResult(Err[int](nil), func(i int) int {
		return i + 1
	})
	assert.True(t, result.IsErr())
}

func TestMapResultOr(t *testing.T) {
	result := MapResultOr(Ok(1), 0, func(i int) int {
		return i + 1
	})
	assert.Equal(t, 2, result)
}

func TestMapResultOr_Err(t *testing.T) {
	result := MapResultOr(Err[int](nil), 0, func(i int) int {
		return i + 1
	})
	assert.Equal(t, 0, result)
}
