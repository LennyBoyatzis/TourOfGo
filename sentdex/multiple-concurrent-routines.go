package main

import "fmt"

func main() {
	out := make(chan int)
	in := make(chan int)

	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	// Go routines do nothing until the get data through in channel
	in <- 1
	in <- 2
	in <- 3

	// Now we wait for each result to come back
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

// IMPORTANT
// there is no guarantee as to which goroutine will accept which input
// or which goroutine will return an output first

func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	num := <-in
	result := num * 2
	out <- result
}
