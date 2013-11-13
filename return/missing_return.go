package main

func main() {
	println(foo())
}

func foo() string {
	return "ok" // valid

	b := true
	if b {
		println("b is true")
	}

	// If you have not the return code below,
	// "missing return at end of function" build error occurred.
	// return "ok"
}
