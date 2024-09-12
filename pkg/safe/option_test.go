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

func TestUnwrap(t *testing.T) {
	s := Some(1)
	assert.Equal(t, 1, s.Unwrap())

	n := None[int]()
	assert.Panics(t, func() {
		_ = n.Unwrap()
	})
}

func TestUnwrapOr(t *testing.T) {
	s := Some(1)
	assert.Equal(t, 1, s.UnwrapOr(0))

	n := None[int]()
	assert.Equal(t, 0, n.UnwrapOr(0))
}
