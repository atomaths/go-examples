// See http://stackoverflow.com/a/11075942

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	producerWG       sync.WaitGroup
	consumerWG       sync.WaitGroup
	resultingChannel = make(chan int)
)

func consumer() {
	defer consumerWG.Done()

	for item := range resultingChannel {
		fmt.Println("Consumed:", item)
	}
}

func producer() {
	defer producerWG.Done()

	success := rand.Float32() > 0.5
	if success {
		resultingChannel <- rand.Int()
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	for c := 0; c < runtime.NumCPU(); c++ {
		producerWG.Add(1)
		go producer()
	}

	for c := 0; c < runtime.NumCPU(); c++ {
		consumerWG.Add(1)
		go consumer()
	}

	producerWG.Wait()

	close(resultingChannel)

	consumerWG.Wait()

	fmt.Println("All done.")
}
