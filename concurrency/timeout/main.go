package main

import (
	"time"
)

func Process() (timeout chan bool) {
	timeout = make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()
	return
}

func main() {
	result := make(chan int)

	// Long running task
	go func() {
		time.Sleep(2 * time.Second)
		result <- 10
	}()

	timeout := Process()

	select {
	case <-result:
	case <-timeout:
	}
}
