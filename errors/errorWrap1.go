package main

import (
	"errors"
	"fmt"
	"os"
)

// have to run with GO version 1.13 or above
func main() {
	var err error
	originalError := &os.PathError{
		Op:   "read",
		Path: "/a/b",
		Err:  errors.New("read error")}

	err = fmt.Errorf("context for error: %v", originalError)
	fmt.Printf("error type: %T, value: %v\n", err, err)

	err = fmt.Errorf("wrapping context for error: %w", originalError)
	fmt.Printf("error type: %T, value: %v\n\n", err, err)

	errInner := errors.Unwrap(err)
	fmt.Printf("error type: %T, value: %v\n", errInner, errInner)
	errInner2 := errors.Unwrap(errInner)
	fmt.Printf("error type: %T, value: %v\n", errInner2, errInner2)

	fmt.Println("Hello, playground")
}
