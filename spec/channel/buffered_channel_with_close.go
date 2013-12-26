// See: https://groups.google.com/forum/#!topic/golang-nuts/7n89ZQfK7F0
//   close builtin function: http://golang.org/ref/spec#Close
//   
// The close source is in $GOROOT/src/pkg/runtime/chan.c:runtime·closechan
//
package main

import "fmt"

func main() {

	ch_1 := make(chan int, 5)

	go Buffered(cap(ch_1), ch_1)
	for i := range ch_1 {
		fmt.Printf("%d ", i)
	}
	// 1 1 2 3

	fmt.Println()
	value_1, ok_1 := <-ch_1
	fmt.Println("MAIN = Buffered value_1:", value_1)
	fmt.Println("MAIN = Buffered ok_1:", ok_1)

	fmt.Println()

	ch_2 := make(chan int)

	go NON_Buffered(5, ch_2)
	for i := range ch_2 {
		fmt.Printf("%d ", i)
	}
	/*
		0 1 1 2 NON_Buffered value: 0
		NON_Buffered ok: false
		3
	*/

	fmt.Println()
	value_2, ok_2 := <-ch_2
	fmt.Println("MAIN = NON_Buffered value_2:", value_2)
	fmt.Println("MAIN = NON_Buffered ok_2:", ok_2)
}

func Buffered(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
	value, ok := <-ch
	fmt.Println("Buffered value:", value)
	fmt.Println("Buffered ok:", ok)
}

func NON_Buffered(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
	value, ok := <-ch
	fmt.Println("NON_Buffered value:", value)
	fmt.Println("NON_Buffered ok:", ok)
}

/*
[Output]
Buffered value: 0
Buffered ok: true
1 1 2 3
MAIN = Buffered value_1: 0
MAIN = Buffered ok_1: false

0 1 1 2 NON_Buffered value: 0
NON_Buffered ok: false
3
MAIN = NON_Buffered value_2: 0
MAIN = NON_Buffered ok_2: false

[Why]
Buffered case is closed at the end
NonBuffered is closed at both cases

---

[Problem #1]
according to Go documentation,
     close(ch)
     value, ok := <-ch
any value received from closed channel is 0
and closed channel sets ok to false
so here the value should be 0
and ok should be 'false', as shown above.

But with buffered channel,
there are some cases that 'value' is 0 and 'ok' is true.

Why Buffered value is 0 which means the closed channel
and ok is TRUE in this case???

And Why Buffered outputs only 4 values?

[Answer #1]
0 here in Buffered is not 0 from the closed channel
'true' means the channel is not closed yet
and 0 is from the initial value.

The first value of five(input) is sent to channel
and gets assigned to 'value_01'
That is why we have only 4 values left to range over.

Since NonBuffered never outputs the TRUE value,
the first value DID NOT get assigned to 'value_02'

If we set the initial value as 10, we get

x, y := 10, 1

Buffered value: 10
Buffered ok: true
1 11 12 23
MAIN = Buffered value_1: 0
MAIN = Buffered ok_1: false

10 1 11 12 NON_Buffered value: 0
NON_Buffered ok: false
23
MAIN = NON_Buffered value_2: 0
MAIN = NON_Buffered ok_2: false



Buffered Channel’s sender does not block
even if the receiver is not ready to receive
unless the buffer is full
, because it has more buffer
so that the receiver can receive the extra in "buffer"

In UnBuffered Channel, by default
sends and receives block until the other side is ready.
That is, every single send will block
until another goroutine receives from the channel.

---

[Problem #2]
Then why the NonBuffered never outputs the TRUE value?
Does this mean the unbuffered channel is always closed?
Then how come the unbuffered channel still produce the right outcome of fibonacci calculation?

[Answer #2]
???

---

[Problem #3]
Why the NonBuffered output gets cut in the middle like the following?

		0 1 1 2 NON_Buffered value: 0
		NON_Buffered ok: false
		3

It should be something related with goroutine but Why in this way?

[Answer #3]
???

---

*/

