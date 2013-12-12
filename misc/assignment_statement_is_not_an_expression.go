// # command-line-arguments
// ./assignment_statement_is_not_an_expression.go:13: pi = nil used as value

package main

import (
	"fmt"
)

func main() {
	i := 5
	pi := &i

	fmt.Printf("%T\n", pi)

	// 다른 언어에서는 if pi = nil 이런 실수하는 것을 막기 위해
	// nil == pi 처럼 쓰는 경우가 있는데,
	// Go 에서는 `assignment statement is not an expression`
	// 이말처럼 nil == pi로 써도 되지만, 그냥 pi == nil 로
	// 쓰기를 권장. pi = nil로 쓴 경우 아예 빌드 에러남.
	if pi == nil {
		println("pi is a pointer")
	} else {
		println("not pointer")
	}
}
