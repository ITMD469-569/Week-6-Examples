// Use a WaitGroup for co-ordination if your program needs to wait for a bunch of Goroutines to finish.
//the counter is decremented using Done.
// taken from https://gobyexample.com/waitgroups

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) { // function that is run every go routine

	defer wg.Done() // decrement waitgroup counter once function finishes

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)         // increment counter for every go routine
		go worker(i, &wg) // launch go routine with waitgroup and i
	}

	wg.Wait() // blocks program until waitgroup counter reaches zero, meaning all go routines are finished
}
