package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root!")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from test!")
}

func main() {
	// declare router, add handler routes, and start the server
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	router.HandleFunc("/test", test)
	http.ListenAndServe(":8000", router)
}
