package main

import (
	"fmt"
)

type Base struct {
	Field1 int
	Field2 string
}

type Derived struct {
	Base   // anonymouse field
	Field3 int
}

// 아래처럼 embedded struct 형태로도 가능
//
//	type Derived struct {
//		Field3 int
//		Base struct {
//			Field1 int
//			Field2 string
//		}
//}

func main() {
	ok1 := Base{3, "Base"}                 // O: untagged struct literal
	ok2 := Base{Field1: 3, Field2: "Base"} // O: tagged struct literal

	// err1 := Base{"Base", 3}             // X: 순서를 바꾸면 안됨
	ok3 := Base{Field2: "Base", Field1: 3} // O: 순서를 바꾸더라도 tagged로 하면 됨

	// err2 := Base{Field1: 3, "Base"}     // X: mixture of field:value and value initializers

	ok4 := Derived{Base{3, "Base"}, 6}

	// err3 := Derived{3, "Base", 6}       // X
	// err4 := Derived{{3, "Base"}, 6}     // X. missing type in composite literal

	// err5 := Derived{6, Base{3, "Base"}} // X: 순서를 바꾸면 안됨
	ok5 := Derived{Field3: 6, Base: Base{3, "Base"}} // O: 순서를 바꾸더라도 tagged로 하면 됨

	fmt.Println(ok1, ok2, ok3, ok4, ok5)
}
