package main

import (
	"fmt"
)

func stringReturn(text string) string {
	return text
}

func main() {
	myText := stringReturn("here is the code")
	fmt.Println(myText)
}
