package safe

import (
	"fmt"
)

type result[T any] struct {
	data T
	err  error
}

// Result is a generic rust-like Result type.
type Result[T any] interface {
	// IsOk returns true if the Result is Ok.
	IsOk() bool
	// IsErr returns true if the Result is Err.
	IsErr() bool
	// Unwrap returns the data if the Result is Ok, panics if the Result is Err.
	Unwrap() T
	// UnwrapFmt returns the data if Result is Ok, panics with a formatted error
	// message if the Result is Err.
	UnwrapFmt(s string) T
	// UnwrapErr returns the error if the Result is Err, panics if the Result
	// is Ok.
	UnwrapErr() error
}

// Ok returns an ok Result.
func Ok[T any](data T) Result[T] {
	return &result[T]{
		data: data,
		err:  nil,
	}
}

// Err returns an errored Result.
func Err[T any](err error) Result[T] {
	if err == nil {
		err = fmt.Errorf("Err called with nil error")
	}

	var t T
	return &result[T]{
		data: t,
		err:  err,
	}
}

// AsResult returns a result from data and an error.  Helpful for converting results
// from normal functions to safe results.
func AsResult[T any](data T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok[T](data)
}

func (r result[T]) IsOk() bool {
	return r.err == nil
}

func (r result[T]) IsErr() bool {
	return r.err != nil
}

func (r result[T]) Unwrap() T {
	if r.IsErr() {
		panic(unwrapError{r.err})
	}
	return r.data
}

func (r result[T]) UnwrapFmt(s string) T {
	if r.IsErr() {
		panic(unwrapError{fmt.Errorf("result is Ok")})
	}
	return r.data
}

func (r result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic(unwrapError{fmt.Errorf("result is Ok")})
	}
	return r.err
}
