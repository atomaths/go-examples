// The return effectively evaluates the return value and stores it before
// any defers execute. But using named parameters will work because they
// provide access to the actual return values.
// See https://groups.google.com/d/msg/golang-nuts/v8tUQHMi4dM/dImbrOFXoHQJ
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := DoNetworkThingWrong()
	fmt.Println("Wrong:", err) // Output: nil
	err = DoNetworkThingRight()
	fmt.Println("Right:", err) // Output: Close error
}

func DoNetworkThingWrong() error {
	conn := Thing{}
	var closeErr error
	defer func() { closeErr = conn.Close() }()

	// use conn

	return closeErr
}

func DoNetworkThingRight() (err error) {
	conn := Thing{}
	defer func() { err = conn.Close() }()

	// use conn

	return
}

type Thing struct{}

func (Thing) Close() error {
	return errors.New("Close error")
}
