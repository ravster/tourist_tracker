package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! You requested: %s\n", r.URL.Path)
}

func main() {
	fmt.Println("Started program")

	http.HandleFunc("/", helloHandler)

	// Start the server on port 8080
	fmt.Println("Starting server at :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
