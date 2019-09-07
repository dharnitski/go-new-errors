package old_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dharnitski/go-new-errors/old"
)

func TestWrapped(t *testing.T) {
	actual := old.Wrapped()
	assert.Equal(t, "wrapper: internal error", actual.Error())
}
