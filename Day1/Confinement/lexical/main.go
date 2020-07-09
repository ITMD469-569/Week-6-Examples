// note channels are concurrent safe, this is just an example of confinement, can be used for other data structures

package main

import (
	"fmt"
)

func main() {

	// writes to the channel
	// confines the write aspect of the channel to the lexical scope of this function
	// this prevents other go routines from working on the channel
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// recieve the channel in read only mode
	// uses range like in the ad hoc example
	// range will read the values are they are recieved until channel is closed (line 15)
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}
