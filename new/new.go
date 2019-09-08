package new

import (
	"errors"
	"fmt"
)

// Wrapped returns wrapped error
func Wrapped() error {
	internal := errors.New("internal error")
	return fmt.Errorf("wrapper: %w", internal)
}

// ErrorsStack wraps error two times
func ErrorsStack() error {
	err := Wrapped()
	return fmt.Errorf("another wrapper: %w", err)
}
