package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)

	t := time.NewTimer(time.Duration(3 * time.Second))

	elapsed := <-t.C
	fmt.Println(elapsed)
}
