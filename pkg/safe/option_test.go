package safe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	s := Some(1)
	assert.True(t, s.IsSome())
	assert.False(t, s.IsNone())
	assert.Equal(t, 1, s.Unwrap())
}

func TestNone(t *testing.T) {
	n := None[int]()
	assert.False(t, n.IsSome())
	assert.True(t, n.IsNone())
	assert.Panics(t, func() {
		_ = n.Unwrap()
	})
}

func TestOption_IsSome(t *testing.T) {
	s := Some(1)
	assert.True(t, s.IsSome())

	n := None[int]()
	assert.False(t, n.IsSome())
}

func TestOption_IsNone(t *testing.T) {
	s := Some(1)
	assert.False(t, s.IsNone())

	n := None[int]()
	assert.True(t, n.IsNone())
}

func TestOption_Expect(t *testing.T) {
	s := Some(1)
	assert.Equal(t, 1, s.Expect("Error"))

	n := None[int]()
	assert.Panics(t, func() {
		_ = n.Expect("Error")
	})
}

func TestOption_Unwrap(t *testing.T) {
	s := Some(1)
	assert.Equal(t, 1, s.Unwrap())

	n := None[int]()
	assert.Panics(t, func() {
		_ = n.Unwrap()
	})
}

func TestOption_UnwrapOr(t *testing.T) {
	s := Some(1)
	assert.Equal(t, 1, s.UnwrapOr(0))

	n := None[int]()
	assert.Equal(t, 0, n.UnwrapOr(0))
}

func TestOption_Ok(t *testing.T) {
	s := Some(1)
	assert.True(t, s.Ok(nil).IsOk())

	n := None[int]()
	assert.True(t, n.Ok(nil).IsErr())
}

func TestOption_String(t *testing.T) {
	s := Some(1)
	assert.Equal(t, "Some(1)", s.String())

	n := None[int]()
	assert.Equal(t, "None", n.String())
}

func TestOption_MarshalJSON(t *testing.T) {
	s := Some(1)
	b, err := s.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "1", string(b))

	n := None[int]()
	b, err = n.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, "{}", string(b))
}

func TestOption_UnmarshalJSON(t *testing.T) {
	var s Option[int]
	err := s.UnmarshalJSON([]byte("1"))
	assert.NoError(t, err)
	assert.True(t, s.IsSome())
	assert.Equal(t, 1, s.Unwrap())

	err = s.UnmarshalJSON([]byte("{}"))
	assert.NoError(t, err)
	assert.True(t, s.IsNone())
}
