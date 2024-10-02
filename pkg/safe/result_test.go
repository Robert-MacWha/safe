package safe

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	result := Ok(1)
	assert.Equal(t, 1, result.Unwrap())
}

func TestErr(t *testing.T) {
	// Test calling with a non-nil error
	result := Err[int](fmt.Errorf("error"))
	assert.Equal(t, "error", result.UnwrapErr().Error())

	// Test calling with a nil error
	result = Err[int](nil)
	assert.Equal(t, "Err called with nil error", result.UnwrapErr().Error())
}

func TestAsResult(t *testing.T) {
	// Test for OK
	result := AsResult(1, nil)
	assert.Equal(t, 1, result.Unwrap())

	// Test for Err
	result = AsResult(1, fmt.Errorf("error"))
	assert.Equal(t, "error", result.UnwrapErr().Error())
}

func TestResult_IsOk(t *testing.T) {
	// Test for OK
	result := Ok(1)
	assert.True(t, result.IsOk())

	// Test for Err
	result = Err[int](fmt.Errorf("error"))
	assert.False(t, result.IsOk())
}

func TestResult_IsErr(t *testing.T) {
	// Test for OK
	result := Ok(1)
	assert.False(t, result.IsErr())

	// Test for Err
	result = Err[int](fmt.Errorf("error"))
	assert.True(t, result.IsErr())
}

func TestResult_Expect(t *testing.T) {
	// Test for OK
	result := Ok(1)
	assert.Equal(t, 1, result.Expect(""))

	// Test for Err
	result = Err[int](fmt.Errorf("error"))
	assert.Panics(t, func() {
		_ = result.Expect("")
	})
}

func TestResult_Unwrap(t *testing.T) {
	// Test for OK
	result := Ok(1)
	assert.Equal(t, 1, result.Unwrap())

	// Test for Err
	result = Err[int](fmt.Errorf("error"))
	assert.Panics(t, func() {
		_ = result.Unwrap()
	})
}

func TestResult_UnwrapOr(t *testing.T) {
	// Test for OK
	result := Ok(1)
	assert.Equal(t, 1, result.UnwrapOr(0))

	// Test for Err
	result = Err[int](fmt.Errorf("error"))
	assert.Equal(t, 0, result.UnwrapOr(0))
}

func TestResult_UnwrapErr(t *testing.T) {
	// Test for OK
	result := Ok(1)
	assert.Panics(t, func() {
		_ = result.UnwrapErr()
	})

	// Test for Err
	result = Err[int](fmt.Errorf("error"))
	assert.Equal(t, "error", result.UnwrapErr().Error())
}

func TestResult_Decompose(t *testing.T) {
	// Test for OK
	result := Ok(1)
	data, err := result.Decompose()
	assert.Equal(t, 1, data)
	assert.Nil(t, err)

	// Test for Err
	result = Err[int](fmt.Errorf("error"))
	data, err = result.Decompose()
	assert.Zero(t, data)
	assert.Equal(t, "error", err.Error())
}

func TestResult_String(t *testing.T) {
	// Test for OK
	result := Ok(1)
	assert.Equal(t, "Ok(1)", result.String())

	// Test for Err
	result = Err[int](fmt.Errorf("error"))
	assert.Equal(t, "Err(error)", result.String())
}
