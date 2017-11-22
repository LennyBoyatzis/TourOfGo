package main

import (
  "fmt"
)

// & - gives you a memory address
// * - gives you the value - reads through memory address 
func main() {
  x := 15
  a := &x // memory address
  fmt.Println("a", a) // 0xc4200160b0
  fmt.Println("a val", *a) // read through memory address
  *a = 5
  fmt.Println(x)
  *a = *a**a
}
