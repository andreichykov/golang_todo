package main

import (
	"encoding/json"
	"fmt"
	"glitch/todo_api/db"
	"glitch/todo_api/dto"
	"io"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func gethandler(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path // Get the path from the request URL
	if requestPath != "/todo" {
		fmt.Printf("Received request for path: %s\n", requestPath)
		io.WriteString(w, fmt.Sprintf("You requested the path: %s\n", requestPath))
		return
	}
	// TODO: @ariel check this out

	obj := dto.Todo{ID: 1, Title: "Learn Go", Completed: true}

	bytes, _ := json.Marshal(obj)
	w.Write(bytes)

}

func main() {
	db.TestDatabase()
	http.HandleFunc("/", gethandler) // Register the handler for all paths
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the server
}
