package safe

import (
	"errors"
	"fmt"
	"reflect"
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
		err = errors.New("Err called with nil error")
	}

	var t T
	return Result[T]{
		data: t,
		err:  err,
	}
}

// As returns a result from data and an error.  Helpful for converting results
// from normal functions to safe results.
func As[T any](data T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(data)
}

// IsOk returns true if the Result is ok.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the Result is an error.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Eq returns true if two Results are equal.
//
// Results are equal if both are ok and their data are DeepEqual, or if both are
// errors and their error messages are equal.
func (r Result[T]) Eq(other Result[T]) bool {
	if r.IsOk() && other.IsOk() {
		return reflect.DeepEqual(r.data, other.data)
	}
	if r.IsErr() && other.IsErr() {
		return r.err.Error() == other.err.Error()
	}
	return false
}

// Expect returns the value of an ok Result or panics with a custom error if the Result is an error.
func (r Result[T]) Expect(msg string) T {
	if r.IsErr() {
		panic(unwrapError{fmt.Errorf(msg+": %w", r.err)})
	}
	return r.data
}

// Unwrap returns the value of an ok Result or panics if the Result is an error.
func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(unwrapError{r.err})
	}
	return r.data
}

// UnwrapOr returns the value of an ok Result or a default value if the Result is an error.
func (r Result[T]) UnwrapOr(def T) T {
	if r.IsErr() {
		return def
	}
	return r.data
}

// UnwrapErr returns the error of an errored Result or panics if the Result is ok.
func (r Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic(unwrapError{errors.New("result is Ok")})
	}
	return r.err
}

// Decompose returns the data and error of a Result.
func (r Result[T]) Decompose() (T, error) {
	return r.data, r.err
}

// String returns a string representation of the Result.
func (r Result[T]) String() string {
	if r.IsOk() {
		return fmt.Sprintf("Ok(%+v)", r.data)
	}
	return fmt.Sprintf("Err(%s)", r.err)
}
