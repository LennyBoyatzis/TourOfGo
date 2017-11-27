package main

import "fmt"

// The chan<- declarations tells us that you can only put stuff
// into the channel, but not receive anything from it

// send-only channel of integers
// out chan<- int

// Receive-only channel of integers
// out <-chan int

// Declare channel without directionality
// out chan int

func multiplyByTwo(num int, out chan<- int) {
	result := num * 2
	out <- result
}

// Channel gives us a way to connect different concurrent paths
// of our program
func main() {
	n := 3
	// Bi-directional channel
	out := make(chan int)

	go multiplyByTwo(n, out)

	// Statements that send or receive values from channels
	// are blocking inside their own goroutine
	fmt.Println(<-out)
	fmt.Println("this wont run until the line above")
}
