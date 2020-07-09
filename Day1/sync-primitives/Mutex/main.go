//A Mutex is a shared lock that you can use to provide exclusive access to certain parts of your code.

package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var functionAccessCount int

func main() {

	var waitGroup sync.WaitGroup
	functionAccessCount = 0 // variable that we will put use the mutex lock on, only accessable when lock is unlocked
	loopNum := 600          // the number of times we want to loop, with mutex, functionAccessCount will be 600 at end
	waitGroup.Add(loopNum)  // add 600 to waitgroup counter

	// start go routine 600 times
	for i := 1; i <= loopNum; i++ {
		go func(c int) {
			defer waitGroup.Done()
			incrementCounter()
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("Final count = ", functionAccessCount)
}

// only a single go routine can access this at once due to the mutex
// showcasing many threads accessing the same function and memory
// when we comment out the mutex, will we get a different result than 600?
func incrementCounter() {
	mutex.Lock()
	defer mutex.Unlock()
	functionAccessCount++
}
