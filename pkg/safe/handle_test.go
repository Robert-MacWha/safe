package safe

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests handle with a panic caused by calling Unwrap on an errored Result
func TestHandle_Result(t *testing.T) {
	var result Result[int]

	assert.NotPanics(t, func() {
		defer Handle(&result)

		// simulate calling unwrap on a err result
		errResult := Err[int](fmt.Errorf("errUnwrap"))
		errResult.Unwrap()
	})

	// Assert that the panic was caught and handled
	assert.False(t, result.IsOk())
	assert.Equal(t, "errUnwrap", result.UnwrapErr().Error())
}

func TestHandle_Option(t *testing.T) {
	var result Result[int]

	assert.NotPanics(t, func() {
		defer Handle(&result)

		// simulate calling unwrap on a err result
		errResult := None[int]()
		errResult.Unwrap()
	})

	// Assert that the panic was caught and handled
	assert.False(t, result.IsOk())
	assert.Equal(t, "called `Unwrap` on `None` value", result.UnwrapErr().Error())
}

// Tests handle with a panic unrelated to the safe package
func TestHandle_Panic(t *testing.T) {
	var result Result[int]

	assert.Panics(t, func() {
		defer Handle(&result)

		// simulate a panic
		panic("panic")
	})
}
