package safe

import (
	"fmt"
)

type Result[T any] struct {
	data T
	err  error
}

// Ok returns an ok Result.
func Ok[T any](data T) Result[T] {
	return Result[T]{
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
	return Result[T]{
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

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(unwrapError{r.err})
	}
	return r.data
}

func (r Result[T]) UnwrapFmt(s string) T {
	if r.IsErr() {
		panic(unwrapError{fmt.Errorf("result is Ok")})
	}
	return r.data
}

func (r Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic(unwrapError{fmt.Errorf("result is Ok")})
	}
	return r.err
}

func (r Result[T]) String() string {
	if r.IsOk() {
		return fmt.Sprintf("Ok(%v)", r.data)
	}
	return fmt.Sprintf("Err(%v)", r.err)
}
