package main

func main() {
	goMain()
}

func goMain() {
	x := 0
scanAgain:
	x++
	if x < 10 {
		goto scanAgain
	}

	println(x)
}
