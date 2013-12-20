package main

import (
	"fmt"
)

type M map[string]interface{}

func insert(docs ...interface{}) error {
	for k, v := range docs {
		fmt.Printf("%v - %v\n", k, v)
	}

	return nil
}

func main() {
	m := M{"key1": 1, "key2": "value2"}
	insert(m)
}
