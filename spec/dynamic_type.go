// Dynamic type with empty interface
package main

import "fmt"

func main() {
	var v interface{} = 2
	fmt.Printf("%v, %T\n", v, v)

	v = "foo"
	fmt.Printf("%v, %T\n", v, v)
}
