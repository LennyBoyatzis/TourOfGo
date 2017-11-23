// channels can be thought of as pipes that connect different concurrent parts of our code

package main

import "fmt"

func main() {
  n := 3
  in := make(chan int)
  out := make(chan int)

  go multiplyByTwo(in, out)

  in <- n
  fmt.Println("What is happening here?")
  output := <-out
  fmt.Println("here is the output", output)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
  fmt.Println("Initializing goroutine...")

  num := <-in

  result := num * 2
  out <- result
}
