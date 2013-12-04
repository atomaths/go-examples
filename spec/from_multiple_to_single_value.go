// http://ernestmicklei.com/2013/11/25/from-multiple-to-single-value-context-in-go/
// https://plus.google.com/+ErnestMicklei/posts/Ktyo1smCRhs
// http://golang.org/ref/spec#Calls

package main

import (
	"fmt"
)

func ab() (a, b int) {
	a, b = 1, 2
	return
}

func first(args ...interface{}) interface{} {
	return args[0]
}

func pick(index int, args ...interface{}) interface{} {
	return args[index]
}

func main() {
	// fmt.Printf("%v\n", ab()) // multiple-value ab() in single-value context error

	fmt.Printf("%v\n", first(ab()))

	// fmt.Printf("%v\n", pick(1, ab())) // multiple-value ab() in single-value context error
	// first()와 같이 ... 파라미터를 가지고 있는데 pick()은 두 번째 args ...interface{}가
	// 받아지지 않는 이유는?
}

// 아래와 같이 first-class function으로 만들 수도 있음.
// pick1(), pick2()로 conversion 하는 형태.
/*
func pair() (int, int) {
	return 1, 2
}

func pick1(i int) func(...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		return args[i]
	}
}

func pick2(args ...interface{}) func(int) interface{} {
	return func(i int) interface{} {
		return args[i]
	}
}

func main() {
	fmt.Println(pick1(1)(pair()))
	fmt.Println(pick2(pair())(1))
}
*/
