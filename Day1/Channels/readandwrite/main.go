// example taken from https://gobyexample.com/channel-directions
// I only added some comments
package main

import "fmt"

// recieves a write only channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// recieves a readonly channel (pings) and a write only channel (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// this example passes a message around channels in different read/write modes
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message") // send message to ping, which writes the messages to the pings channel in write only mode
	pong(pings, pongs)            // recieves the pings channel in read only mode, writes message in pings channel into pongs channel, which is in write only mode
	fmt.Println(<-pongs)          // print out the message in the pongs channel
}
