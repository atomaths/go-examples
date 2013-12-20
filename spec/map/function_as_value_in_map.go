// To use function as value and integer as key in map.
// And the function can called.
package main

import "fmt"

func main() {
	funcmap := map[int]func() int{1: func() int {
		return 10
	}, 2: func() int {
		return 20
	}}
	fmt.Printf("%T => %v\n", funcmap[1], funcmap[1])
	fmt.Printf("%T => %v\n", funcmap[1](), funcmap[1]())
}
