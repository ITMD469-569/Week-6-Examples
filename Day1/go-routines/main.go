package main

import (
	"fmt"
	"time"
)

func doInBackground() {
	for i := 0; i < 10; i++ {
		percent := i * 10
		fmt.Println("Background task progress:", percent, "%")
		time.Sleep(time.Second)
	}
}

func main() {
	go doInBackground()
	for i := 0; i < 10; i++ {
		percent := i * 10
		fmt.Println("Foreground task progress:", percent, "%")
		time.Sleep(time.Second * 2)
	}

}
