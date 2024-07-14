package errors

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

// New is a convenience wrapper around errors.New.
func New(msg string) error {
	return errors.New(msg)
}

// Is is a convenience wrapper around errors.Is.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As is a convenience wrapper around errors.As.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Wrap wraps an error with the file and line where it was called.
func Wrap(err error) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("unknown caller: %w", err)
	}
	srcFile := filepath.Base(file)
	dir := filepath.Base(filepath.Dir(file))
	file = filepath.Join(dir, srcFile)
	return fmt.Errorf("%s:%d: %w", file, line, err)
}

// Unwrap is a convenience wrapper around errors.Unwrap.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
