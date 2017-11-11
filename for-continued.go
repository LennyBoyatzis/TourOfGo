package main

import "fmt"

func forWithOutStatements() {
  sum := 1
  for ; sum < 1000 ; {
    sum += sum
  }
  fmt.Println(sum)
}

func main() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}
