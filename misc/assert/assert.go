package main

func assert(cond bool, msg string) {
	if !cond {
		panic("[package name] internal error: " + msg)
	}
}

func assertZero(i int) {
	assert(i == 0, "zero")
}

func main() {
	assertZero(0)
}
