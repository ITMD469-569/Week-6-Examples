package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from root!\n")
	fmt.Fprintf(w, "Request URL path: "+r.URL.String())
	fmt.Fprintf(w, "Request URL Body: ", r.Body)
	fmt.Fprintf(w, "Request URL header: ", r.Header)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from test!")
	fmt.Fprintf(w, "Request URL path: "+r.URL.String())
	fmt.Fprintf(w, "Request URL Body: ", r.Body)
	fmt.Fprintf(w, "Request URL header: ", r.Header)
}

func main() {
	// declare router, add handler routes, and start the server
	router := mux.NewRouter()

	router.HandleFunc("/", handler)
	router.HandleFunc("/test", test)

	http.ListenAndServe(":8000", router)
}
