package main

import (
	"encoding/json"
	"fmt"
	"glitch/todo_api/dto"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// requestPath := r.URL.Path // Get the path from the request URL

	// NOT GET
	// if requestPath != "/todo" {
	// 	fmt.Printf("Received request for path: %s\n", requestPath)
	// 	io.WriteString(w, fmt.Sprintf("You requested the path: %s\n", requestPath))
	// 	return
	// }

	// NOT GET

	// obj := dto.Todo{ID: 1, Title: "Learn Go", Completed: true}

	bytes, _ := json.Marshal(dto.Todos)
	w.Write(bytes)

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var newTodo dto.TodoPost
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newTodo.ID = dto.IDTodo
	dto.IDTodo++
	dto.Todos = append(dto.Todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func main() {
	// db.TestDatabase()
	http.HandleFunc("/todo", PostHandler)
	http.HandleFunc("/", GetHandler) // Register the handler for all paths
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the server
}
