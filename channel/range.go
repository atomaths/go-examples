package main

import (
	"time"
)

/*
type Worker struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		time.Sleep(w.z * time.Second)
		out <- w
	}
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < NumWorkers; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
}

func main() {
	Run()
}
*/

func worker(in <-chan int) {
	println("start worker")
	for w := range in {
		println("received data: ", w)
	}
	println("end worker")
}

func main() {
	in := make(chan int)
	go worker(in)
	go worker(in)
	go worker(in)

	time.Sleep(3 * time.Second)
	in <- 10
	time.Sleep(3 * time.Second)
}
