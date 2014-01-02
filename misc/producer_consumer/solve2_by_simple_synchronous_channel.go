// See http://stackoverflow.com/questions/11075876/what-is-the-neatest-idiom-for-producer-consumer-in-go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rand.Float32() 결과가 0.5 보다 크면 rand.Int() 값을 
// consumer가 출력할 수 있도록 c 채널로 보냄. 
// 0.5보다 크지 않다면 consumer에서 출력할 수 있는 개수가
// 그만큼 줄어들게 됨(총 nTries 만큼 done 채널에서 받도록
// main 함수에서 되어 있음). fail 채널로 false를 보내기 때문에.
func producer(c chan<- int, fail chan<- bool) {
	if success := rand.Float32() > 0.5; success {
		c <- rand.Int()
	} else {
		fail <- true
	}
}

func consumer(c <-chan int, success chan<- bool) {
	for {
		num := <-c
		fmt.Println("Consumed:", num)
		success <- true
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	const nTries = 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < nTries; i++ {
		go producer(c, done)
	}
	go consumer(c, done)

	for i := 0; i < nTries; i++ {
		<-done
	}

	fmt.Println("All done.")
}
