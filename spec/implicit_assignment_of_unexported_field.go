// package mypkg
//
// type MyPkg struct {
//	string
// }

package main

import "mypkg"

func main() {
	// Build error:
	// implicit assignment of unexported field 'string' in un.MyError literal
	_ := mypkg.MyPkg{"foo"}
}
