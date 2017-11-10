package main

import (
  "fmt"
  "io/ioutil"
  m "math"
  "net/http"
  "os"
  "strconv"
)

func main() {
  fmt.Println("Hello world!")
  beyondHello()
}

// Declarations use := to infer the type

func beyondHello() {
  var x int
  x = 3
  y := 4
  sum, prod := learnMultiple(x, y)
  fmt.Println("sum: ", sum, "prod: ", prod)
  learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
  return x + y, x * y // Return two vals
}

func learnTypes() {
  str := "Learn Go!"
  s2 := `A "raw" string
  literal can include line breaks`

  g := 'Σ'
  f := 3.14195
  c := 3 + 4i

  var u uint = 7
  var pi float32 = 22. / 7

  n := byte('\n')
}
