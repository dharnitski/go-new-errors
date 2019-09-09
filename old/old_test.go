package old_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dharnitski/go-new-errors/new"
	"github.com/dharnitski/go-new-errors/old"
)

// TestWrapped 
func TestWrapped(t *testing.T) {
	actual := old.Wrapped()
	assert.Equal(t, "wrapper: internal error", actual.Error())
}

func TestUnwrap(t *testing.T) {
	wrapped := old.Wrapped()
	unwrapped := errors.Cause(wrapped)
	require.NotNil(t, unwrapped)
	assert.Equal(t, "internal error", unwrapped.Error())
}

func TestStack(t *testing.T) {
	stack := old.ErrorsStack()
	assert.Equal(t, "another wrapper: wrapper: internal error", stack.Error())

	one := errors.Cause(stack)
	require.NotNil(t, one)
	assert.Equal(t, "internal error", one.Error())

	two := errors.Cause(one)
	require.NotNil(t, two)
	assert.Equal(t, "internal error", two.Error())
}

func TestFormat(t *testing.T) {
	err := old.Wrapped()
	actual := fmt.Sprintf("%+v\n", err)
	assert.Contains(t, actual, "internal error")
	assert.Contains(t, actual, "github.com/dharnitski/go-new-errors/old.Wrapped")
	assert.Contains(t, actual, "/github.com/dharnitski/go-new-errors/old/old.go:")
	assert.Contains(t, actual, "wrapper")
}

func TestUnwrapNew(t *testing.T) {
	wrapped := new.Wrapped()
	unwrapped := errors.Cause(wrapped)
	require.NotNil(t, unwrapped)
	assert.Equal(t, "wrapper: internal error", unwrapped.Error())
}
