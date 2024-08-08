package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a simple handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Start the server on port 8080
	fmt.Println("Server is running on port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
