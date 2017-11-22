package main

import "fmt"

// defer statements are done in a first in last out order
func foo() {
  defer fmt.Println("Done!")
  fmt.Println("Doing some stuff, who knows what?")
  defer fmt.Println("Then Im done")
  defer fmt.Println("And now Then Im done")
}

// go routine is a lightweight there
// if your program finishes before the goroutine finishes then its over
func main() {
  foo()
}
