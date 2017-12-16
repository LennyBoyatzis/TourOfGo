package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, goGroup *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	goGroup.Done()
}

func main() {

	// Declare a WaitGroup struct
	goGroup := new(sync.WaitGroup)
	fmt.Println("Starting")

	hello := new(Job)
	world := new(Job)

	hello.text = "hello"
	hello.i = 0
	hello.max = 5

	world.text = "world"
	world.i = 0
	world.max = 5

	// When a goroutine is invoked,
	// it waits for blocking code to complete
	// before concurrency begins

	go outputText(hello, goGroup)
	go outputText(world, goGroup)

	// Number of Done messages should receive before satisfying wait
	goGroup.Add(2)
	goGroup.Wait()

	fmt.Println("Finishing")
}
