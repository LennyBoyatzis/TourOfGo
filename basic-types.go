package main

// BASIC TYPES
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte (uint8)
// rune (int32)
// float32 float64
// complex64 complex128

import (
  "fmt"
  "math/cmplx"
)

// variable declarations may be factored into blocks
// When you have a int
// You should use int unless you have a specific reason
// to use a sized or unsigned integer type
var (
  ToBe bool = false
  MaxInt uint64 = 1<<64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)
}


