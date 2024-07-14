package errors_test

import (
	"io/fs"
	"os"
	"testing"

	"dario.cat/errors"
)

func TestNew(t *testing.T) {
	t.Parallel()
	err := errors.New("test")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "test" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestIs(t *testing.T) {
	t.Parallel()
	_, err := os.Open("non-existent-file")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	err = errors.Wrap(err)
	if !errors.Is(err, fs.ErrNotExist) {
		t.Fatal("err should be fs.ErrNotExist")
	}
}

func TestAs(t *testing.T) {
	t.Parallel()
	_, err := os.Open("non-existent-file")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	err = errors.Wrap(err)
	var pathError *fs.PathError
	if !errors.As(err, &pathError) {
		t.Fatal("err should include a *fs.PathError")
	}
}

func TestWrap(t *testing.T) {
	t.Parallel()
	err := errors.New("test")
	err = errors.Wrap(err)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "errors/errors_test.go:12: test" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUnwrap(t *testing.T) {
	t.Parallel()
	err := errors.New("test")
	err = errors.Wrap(err)
	err = errors.Unwrap(err)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "test" {
		t.Fatalf("unexpected error: %v", err)
	}
}
