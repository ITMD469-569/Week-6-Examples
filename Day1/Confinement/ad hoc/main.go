// achieve confinement through a convention
// this works well, but can be hard to notice by large teams
// lexical confinement is preferred

package main

import "fmt"

func main() {
	data := make([]int, 4)

	handleData := make(chan int)  // retrieves values from data in loopData
	go loopData(handleData, data) // in the background, populate handleData with data

	// print out items in handleData
	counter := 0

	// range iterates over each element in the channel as it is RECIEVED
	// once the channel is closed, the range function will stop execution
	for num := range handleData {
		fmt.Println(num)
		counter++
	}

	fmt.Println("Number of elements in handleData: ", counter)
}

// pass an entire copy of the data slice to function to operate on
// channel is in write only mode, cannot be read from
func loopData(handleData chan<- int, data []int) {
	defer close(handleData) // close the channel once we are done, causing range to stop executing
	for i := range data {
		handleData <- data[i] // put items from data into handleData channel
	}
}
