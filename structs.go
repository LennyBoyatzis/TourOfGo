package main

import "fmt"

// A struct is a collection of fields

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X)
}
