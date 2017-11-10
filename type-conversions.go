package main

import (
  "fmt"
  "math"
)

// In go an assignment between items of different type
// requires an explicit conversion

func main() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)
}
