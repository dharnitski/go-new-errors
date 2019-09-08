package new_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dharnitski/go-new-errors/new"
)

func TestWrapped(t *testing.T) {
	actual := new.Wrapped()
	assert.Equal(t, "wrapper: internal error", actual.Error())
}

func TestUnwrap(t *testing.T) {
	wrapped := new.Wrapped()
	unwrapped := errors.Unwrap(wrapped)
	require.NotNil(t, unwrapped)
	assert.Equal(t, "internal error", unwrapped.Error())
}

func TestStack(t *testing.T) {
	stack := new.ErrorsStack()
	assert.Equal(t, "another wrapper: wrapper: internal error", stack.Error())

	one := errors.Unwrap(stack)
	require.NotNil(t, one)
	// this is different, old returns internal error
	assert.Equal(t, "wrapper: internal error", one.Error())

	two := errors.Unwrap(one)
	require.NotNil(t, two)
	assert.Equal(t, "internal error", two.Error())
}

func TestFormat(t *testing.T) {
	err := new.Wrapped()
	actual := fmt.Sprintf("%+v\n", err)
	// difference: github.com/pkg/errors returns callstack for all nested errors for +v
	assert.Equal(t, "wrapper: internal error\n", actual)
}
