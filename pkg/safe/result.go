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
	// UnwrapErr returns the error if the Result is Err, panics if the Result
	// is Ok.
	UnwrapErr() error
	// UnwrapFmt returns a formatted error if the Result is Err, panics if the
	// Result is Ok.
	UnwrapFmt(s string) error
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

func Res[T any](data T, err error) Result[T] {
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

func (r result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic(unwrapError{fmt.Errorf("result is Ok")})
	}
	return r.err
}

func (r result[T]) UnwrapFmt(s string) error {
	if r.IsOk() {
		panic(unwrapError{fmt.Errorf("result is Ok")})
	}
	return fmt.Errorf(s, r.err)
}
