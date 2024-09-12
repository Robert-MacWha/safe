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

func (o Option[T]) IsSome() bool {
	return o.Value != nil
}

func (o Option[T]) IsNone() bool {
	return o.Value == nil
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic(unwrapError{fmt.Errorf("attempted to unwrap None")})
	}
	return *o.Value
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.IsNone() {
		return def
	}
	return *o.Value
}

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
