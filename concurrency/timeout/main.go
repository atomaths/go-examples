package main

import (
	"time"
)

func TimeCheck() (timeout chan bool) {
	timeout = make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()
	return
}

func main() {
	ch := make(chan int)

	// Long running task
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 10
	}()

	timeout := TimeCheck()

	select {
	case result := <-ch:
		println("result received: ", result)
	case <-timeout:
		println("timeout")
	}
}
