// the default server mux is limited and does not have great preformance
// Getting information from the requests can be cumbersome

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root!")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from test!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
