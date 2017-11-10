package main

import "fmt"

func add(a int, b int) int {
  return a + b
}

func runForLoop() {
  sum := 0
  for i := 0; i < 8; i++ {
    sum += i
    fmt.Println("sum", sum)
  }
}

func primes() {
  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)

  var s []int = primes[1:4]
  fmt.Println(s)
}

type Vertex struct {
  X int
  Y int
}

func testVertex() {
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X)
}

func main() {
  runForLoop()
  primes()
  testVertex()
}
