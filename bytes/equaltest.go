package main

import (
	"fmt"
	"bytes"
)

func main() {
	b1 := []byte{'a', 'b', 'c'}
	b2 := []byte{'a', 'b', 'c'}
	fmt.Printf("%t\n", bytes.Equal(b1, b2)) // true

	b3 := []byte("abc") // type conversion
	b4 := []byte("abc")
	fmt.Printf("%t\n", bytes.Equal(b3, b4)) // true

	println(b1[0], b3[0]) // 97, 97


	b5 := []byte("한글")
	b6 := []byte("한글")
	fmt.Printf("%t\n", bytes.Equal(b5, b6)) // true
}
