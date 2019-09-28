package main

import (
	"fmt"
)

type BusinessError struct {
	err string
}

// any struct that has Error() method will meet error interface
func (b *BusinessError) Error() string {
	return b.err
}

func main() {
	errBiz := &BusinessError{}
	errBiz.err = "Testing"

	// Println check and find error type will call Error() to print out
	fmt.Println(errBiz)
}
