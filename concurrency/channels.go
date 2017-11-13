package main

import "fmt"

func sum(s []int, c chan int) {
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum // send sum to c
}

// By default sends and receives will block until the other
// side is ready

func main() {
  s := []int{7, 2, 8, -9, 4, 0}

  c := make(chan int)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)
  // once both go routines have completed their computation
  // it calculates the final result
  x, y := <-c, <-c

  fmt.Println(x, y, x+y)
}
