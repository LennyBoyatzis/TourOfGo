package main

// Idea with channels
// Use them with go routines in order
// to send and receive values between them
import "fmt"

func foo(c chan int, someValue int) {
  c <- someValue * 5
}

func main() {
  fooVal := make(chan int)

  go foo(fooVal, 5)
  go foo(fooVal, 3)

  val1 := <-fooVal
  val2 := <-fooVal

  fmt.Println(v1, v2)
}
