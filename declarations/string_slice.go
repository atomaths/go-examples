package main

import "fmt"

func main() {
	fmt.Printf("%T\n", []string{"abc"}) // []string
	fmt.Printf("%T\n", string("abc")) // string
	fmt.Printf("%T\n", "abc") // string
}
