package old

import "github.com/pkg/errors"

// Wrapped returns wrapped error
func Wrapped() error {
	internal := errors.New("internal error")
	return errors.Wrap(internal, "wrapper")
}

// ErrorsStack wraps error two times
func ErrorsStack() error {
	err := Wrapped()
	return errors.Wrap(err, "another wrapper")
}
