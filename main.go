package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// Associate the root path "/" with our helloHandler function
	http.HandleFunc("/", helloHandler)

	// 1. Read the port provided dynamically by Vercel
	port := os.Getenv("PORT")

	// 2. Fallback to 3000 (or any port) for local development
	if port == "" {
		port = ":3000"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	// Start the server on port 5004
	fmt.Println("Server is running on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
