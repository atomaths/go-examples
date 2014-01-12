// trace가 리턴하는 값을 un의 인자로 받아 처리
// struct 같은 걸로 좀더 복잡한 것도 할 수 있다.
package main

import "fmt"

func main() {
	forMain()
}

func forMain() {
	defer un(trace(1))

	fmt.Println("    do something...")
}

func trace(i int) int {
	fmt.Println(i, "start (")
	return i
}

func un(i int) {
	fmt.Println(") end ", i)
}
