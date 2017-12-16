package main

// Defer is a way of delaying the execution of a
// statement until the rest of the function is complete

import (
	"fmt"
)

func main() {
	aValue := new(int)
	defer fmt.Println(*aValue)

	for i := 0; i < 100; i++ {
		*aValue++
	}
}
