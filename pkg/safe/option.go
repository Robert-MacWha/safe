package safe

import (
	"encoding/json"
	"fmt"
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

// Expect returns the value of a Some Option or panics with a custom error if the Option is a None.
func (o Option[T]) Expect(err error) T {
	if o.IsNone() {
		panic(unwrapError{err})
	}
	return *o.Value
}

// Unwrap returns the value of a Some Option or panics if the Option is a None.
func (o Option[T]) Unwrap() T {
	return o.Expect(fmt.Errorf("called `Unwrap` on a `None` value"))
}

// UnwrapOr returns the value of a Some Option or a default value if the Option is a None.
func (o Option[T]) UnwrapOr(def T) T {
	if o.IsNone() {
		return def
	}
	return *o.Value
}

// Map applies a function to the value of a Some Option and returns a new Option.
func (o Option[T]) Map(f func(T) T) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	return Some(f(*o.Value))
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
	return fmt.Sprintf("Some(%s)", *o.Value)
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
