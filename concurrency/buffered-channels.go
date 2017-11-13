package main

import "fmt"

// Channels can be buffered.
// If you provide the buffer length as a second argument to make
// to initialize a buffered channel

// Will send to a buffered channel block only when
// the buffer is full. Receives a block when buffer
// is empty

func main() {
  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  fmt.Println(<-ch)
  fmt.Println(<-ch)
}
