package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// Associate the root path "/" with our helloHandler function
	http.HandleFunc("/", helloHandler)

	// Start the server on port 5004
	fmt.Println("Server is running on http://localhost:5004")
	err := http.ListenAndServe(":5004", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
