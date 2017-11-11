package main

import "fmt"

// Go has only one looping construct
// The for loop
// init statement
// condition expression
// post statement

func main() {
  sum := 0
  for i:= 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
