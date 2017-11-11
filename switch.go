package main

import (
  "fmt"
  "runtime"
)

// Go only runs the selected case, not all the cases that follow
// In effect, the break statement is automatically provided

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux. ")
  default:
    fmt.Println("%s.", os)
  }
}
