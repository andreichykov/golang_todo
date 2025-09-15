package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path // Get the path from the request URL
	fmt.Printf("Received request for path: %s\n", requestPath)
	io.WriteString(w, fmt.Sprintf("You requested the path: %s\n", requestPath))
}

func main() {
	http.HandleFunc("/", handler) // Register the handler for all paths
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the server
}
