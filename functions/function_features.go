package main

import (
	"fmt"
)

func MultipleReturn(i, d int) (int, int) {
	return i / d, i % d
}

type userDefinedFunc func(i int) int


type Request struct {
	Uri string
	RequestTime int64
}

func main() {
	// Multiple return values
	quotient, remainder := MultipleReturn(9, 2)
	fmt.Println(quotient, remainder)
	// Output:
	// 4, 1

	// User defined function types
}
