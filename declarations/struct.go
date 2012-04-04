package main

import (
	"fmt"
)

/*
type Request struct {
	Protocol string
	Time int64
	Foo, Bar string
}

var Requests = []Request{
	{"HTTP", 1234, "Foo", "Bar"},
	{"SPDY", 1234, "A", "B"},
}
*/

// 위처럼, Request struct 타입을 만들고 그 Request 타입으로 literal을 해도 되지만
// 아래처럼 이름이 없는 struct 형태를 선언과 동시에 literal로 변수를 만들 수도 있음
var Request = []struct {
	Protocol string
	Time int64
	Foo, Bar string
}{
	{"HTTP", 1234, "Foo", "Bar"},
	{"SPDY", 1234, "A", "B"},
}

func main() {
	for k, v := range Request {
		fmt.Println(k, v)
	}
}
