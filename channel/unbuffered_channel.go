package main

func main() {
	ch := make(chan bool)
	i := 10

	go func() {
		println(i)
		ch <- true
	}()

	<-ch // A receive from an unbuffered channel happens before the send on that channel completes.
}
