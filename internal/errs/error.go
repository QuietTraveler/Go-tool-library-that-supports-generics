// Package errs error.go,Store error messages
package errs

import "fmt"

// NewErrIndexOutOfRange Create subscript out of range error
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("gtool: Subscript is not in range, The length:%d, index:%d", length, index)
}

// NewErrInvalidType Create type mismatch error
func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("gtool: Type conversion failed, expected type:%s, actual value:%#v", want, got)
}
