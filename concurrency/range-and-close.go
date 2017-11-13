package main

import "fmt"

// sender can close channel to indicate
// no more vals will be sent

// Revceivers can test whether a channel
// has been closes by assigning a second param

// v, ok := <=ch

func fibonacci(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  close(c)
}

func main() {
  c := make(chan int, 10)
  go fibonacci(cap(c), c)
  for i := range c {
    fmt.Println(i)
  }
}
