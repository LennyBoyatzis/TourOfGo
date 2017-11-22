package main

import (
  "time"
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func cleanup() {
  if r := recover(); r != nil {
    fmt.Println("Recovered in cleanup", r)
  }
  wg.Done()
}

func say(s string) {
  defer cleanup()
  for i := 0; i < 3; i++ {
    time.Sleep(time.Millisecond*100)
    fmt.Println(s)
    if i == 2 {
      panic("Oh no")
    }
  }
}

// go routine is a lightweight thread
// if your program finishes before the goroutine finishes then its over
func main() {
  wg.Add(1)
  go say("hey")
  wg.Add(1)
  go say("there")
  wg.Wait()
}
