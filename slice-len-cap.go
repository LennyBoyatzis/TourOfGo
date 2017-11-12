package main

import "fmt"

// A slice has both a length and a capacity
// Length of the slice is the number of elements it contains
// The capacity of the slice is the number of elements in the underlying array
// You can extend a slices length by re-slicing it, provided it has sufficient
// capacity

func main() {
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)

  // Slice the slice to give it a zero val
  s = s[:0]
  printSlice(s)

  // Extend its length
  s = s[:4]
  printSlice(s)

  // Drop its first two vals
  s = s[2:]
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
