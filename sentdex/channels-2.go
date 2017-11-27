// channels can be thought of as pipes that connect different concurrent parts of our code

package main

import "fmt"

func main() {
	n := 3
	in := make(chan int)
	out := make(chan int)

	go multiplyByTwo(in, out)

	in <- n
	output := <-out
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	num := <-in

	result := num * 2
	out <- result
}
