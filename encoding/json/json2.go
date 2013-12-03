package main

import (
	"fmt"
	"encoding/json"
)

type Base struct {
	FieldA int
	FieldB string
}

type Derived struct {
	Base
	FieldC int
}

func main() {
	d := Derived{Base{3,"base"}, 5}
	fmt.Printf("%v\n", d.Base.FieldB)

	b, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", b.Base.FieldA)
}
