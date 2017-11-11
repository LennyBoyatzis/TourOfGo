package main

import (
  "fmt"
  "math"
)
// if statements can start with a short statement to execute
// before the condition
// Only in scope until the end of the if

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
