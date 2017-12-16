package main

// When an application or process can manage its own thread
// or scheduling, it results in a faster runtime
// Threads granted to our app and Go's scheduler have fewer OS
// attributes that need to be considered in context to switching

// Go uses its own scheduler
// If you don't yield to the main thread
// Your goroutines will perform in an unexpected way

// A go routine must be blocked before concurrency is valid
// and can begin

// Since the Go scheduler managers context switching,
// each goroutine must yield control back to the main thread
// to schedule all async tasks

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job) {
	fileName := j.text + ".txt"
	fileContents := ""

	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fileContents += j.text
		fmt.Println(j.text)
		j.i++
	}
	err := ioutil.WriteFile(fileName, []byte(fileContents), 0644)
	if err != nil {
		panic("Something went wrong")
	}
}

func main() {
	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 5

	go outputText(hello)
	go outputText(world)

	runtime.Gosched()
}
