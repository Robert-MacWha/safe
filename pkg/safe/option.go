package safe

import "fmt"

// Option is a generic rust-like Option type.
type Option[T any] interface {
	// IsSome returns true if the Option is Some.
	IsSome() bool
	// IsNone returns true if the Option is None.
	IsNone() bool
	// Unwrap returns the data if the Option is Some, panics if the Option is None.
	Unwrap() T
	// UnwrapOr returns the data if the Option is Some, otherwise returns the
	// default value.
	UnwrapOr(def T) T
}

type option[T any] struct {
	Value *T `json:"value,omitempty" bson:"value,omitempty"`
}

// Some returns a Some Option.
func Some[T any](data T) Option[T] {
	return &option[T]{
		Value: &data,
	}
}

// None returns a None Option.
func None[T any]() Option[T] {
	return &option[T]{
		Value: nil,
	}
}

func (o option[T]) IsSome() bool {
	return o.Value != nil
}

func (o option[T]) IsNone() bool {
	return o.Value == nil
}

func (o option[T]) Unwrap() T {
	if o.IsNone() {
		panic(unwrapError{fmt.Errorf("attempted to unwrap None")})
	}
	return *o.Value
}

func (o option[T]) UnwrapOr(def T) T {
	if o.IsNone() {
		return def
	}
	return *o.Value
}
