# errors

`errors` package is a wrapper around Go's `errors` package. It provides a way to wrap errors with additional context about where the error occurred.

It's a simple package to avoid the need to write the following boilerplate to make errors easier to track.

```go
if _, err := os.Open("non-existent-file"); err != nil {
	return fmt.Errorf("os.Remove: %w", err)
}
```

Instead, you can write:

```go
if _, err := os.Open("non-existent-file"); err != nil {
	return errors.Wrap(err)
}
```

The error will be wrapped with the file and line number where the error occurred. `Error` will return a string like:

```
example.go:17: file does not exist
```
