package main

import "fmt"

// Struct fields can be accessed through a struct pointer
// To access the field X of a struct through a pointer
// we could write (*p).X - shorthand is just p.X

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
