package main

import (
  "time"
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func say(s string) {
  defer wg.Done()
  for i := 0; i < 3; i++ {
    fmt.Println(s)
    time.Sleep(time.Millisecond*100)
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
