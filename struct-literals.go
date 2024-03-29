package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1, 2} // has type Vertex
  v2 = Vertex{X: 1} // Y:0 is implicit
  v3 = Vertex{}     // v3 = Vertex{}
  p = &Vertex{1, 2} // p = &Vertex{1, 2}
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
