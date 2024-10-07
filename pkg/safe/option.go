package safe

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Option[T any] struct {
	Value *T `json:"value,omitempty" bson:"value,omitempty" rethinkdb:"value,omitempty"`
}

// Some returns a Some Option.
func Some[T any](data T) Option[T] {
	return Option[T]{
		Value: &data,
	}
}

// None returns a None Option.
func None[T any]() Option[T] {
	return Option[T]{
		Value: nil,
	}
}

// IsSome returns true if the Option is a Some.
func (o Option[T]) IsSome() bool {
	return o.Value != nil
}

// IsNone returns true if the Option is a None.
func (o Option[T]) IsNone() bool {
	return o.Value == nil
}

// Eq returns true if two Options are equal.
//
// Options are equal if both are Some and their values are DeepEqual, or if both are None.
func (o Option[T]) Eq(other Option[T]) bool {
	if o.IsNone() && other.IsNone() {
		return true
	}
	if o.IsSome() && other.IsSome() {
		return reflect.DeepEqual(o.Value, other.Value)
	}
	return false
}

// Expect returns the value of a Some Option or panics with a custom error if the Option is a None.
// Expects printf-style arguments.
func (o Option[T]) Expect(msg string, a ...any) T {
	msg = fmt.Sprintf(msg, a...)
	if o.IsNone() {
		panic(unwrapError{errors.New(msg)})
	}
	return *o.Value
}

// Unwrap returns the value of a Some Option or panics if the Option is a None.
func (o Option[T]) Unwrap() T {
	return o.Expect("called `Unwrap` on `None` value")
}

// UnwrapOr returns the value of a Some Option or a default value if the Option is a None.
func (o Option[T]) UnwrapOr(def T) T {
	if o.IsNone() {
		return def
	}
	return *o.Value
}

func (o Option[T]) Decompose() (T, bool) {
	if o.IsNone() {
		return *new(T), false
	}
	return *o.Value, true
}

// Ok returns a Result containing the value of a Some Option or an error if the Option is a None.
func (o Option[T]) Ok(err error) Result[T] {
	if o.IsNone() {
		return Err[T](err)
	}
	return Ok(*o.Value)
}

// String returns a string representation of the Option.
func (o Option[T]) String() string {
	if o.IsNone() {
		return "None"
	}
	return fmt.Sprintf("Some(%v)", *o.Value)
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsNone() {
		return []byte("{}"), nil
	}
	return json.Marshal(o.Value)
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		if strings.HasPrefix(string(data), "{}") {
			o.Value = nil
			return nil
		}
		return err
	}
	o.Value = &result
	return nil
}
