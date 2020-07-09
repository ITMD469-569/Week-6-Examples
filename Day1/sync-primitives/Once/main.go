//Once allows you to define a task which you only want to execute once during the lifetime of your program.

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var once sync.Once

func main() {
	// declare http handler for root
	// only execute onlyDoOnce the first time the handler it hit
	// will execute the other functions per hit
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("http handler start")
		once.Do(onlyDoOnce)
		fmt.Println("http handler end")
		w.Write([]byte("done!"))
	})

	http.ListenAndServe(":8080", nil)
}

func onlyDoOnce() {
	fmt.Println("This function will only run once, even though handleFunc can be called many times")
}
