package main

import (
  "fmt"
  "time"
)

// A goroutine is a lightweight thread managed by the go runtime
// goroutines run in the same address space, access to shared memory must be synchronised

func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func main() {
  go say("world")
  say("Hello")
}
