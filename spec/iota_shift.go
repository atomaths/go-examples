package main

import "fmt"

const (
	A = iota
	B
	C
)

const (
	AA = 2 << iota
	BB
	CC
)

const (
	DEBUG = 1 << iota
	INFO
	WARN
	ERROR
	FATAL
)

type Mode uint

const (
	RawFormat Mode = 1 << iota
	TabIndent
	UseSpaces
	SourcePos
)

func main() {
	// 2<<10 == 2^10

	fmt.Println(A, B, C)                         // Output: 0 1 2
	fmt.Println(AA, BB, CC)                      // Output: 2 4 8
	fmt.Println(DEBUG, INFO, WARN, ERROR, FATAL) // Output: 1 2 4 8 16

	var mode Mode
	mode = UseSpaces | TabIndent
	fmt.Println(mode) // Output: 6
}
