package main

import "fmt"

// deferred func calls are push onto a stack
// once the surrounding func has returned
// calls are executed in a last-in-first-out order

func main() {
  fmt.Println("counting")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("done")
}
