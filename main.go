package main

import (
	"encoding/json"
	"fmt"
	"glitch/todo_api/dto"
	"io"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func Handler(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path // Get the path from the request URL

	// Print the request path
	if requestPath != "/todo" {
		fmt.Printf("Received request for path: %s\n", requestPath)
		io.WriteString(w, fmt.Sprintf("You requested the path: %s\n", requestPath))
		return
	}

	bytes, _ := json.Marshal(dto.Todos)
	w.Write(bytes)

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo dto.TodoPost
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
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

func GetHandler(w http.ResponseWriter, r *http.Request) {
	TodoList := dto.GetTodos()
	if r.Method == http.MethodGet {

		bytes, _ := json.Marshal(TodoList)
		w.Write(bytes) // get method all todos
	}

}

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/GetTodo/"))
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}

	for _, todo := range dto.Todos {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound) // get todo by id
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		dto.Todos = nil

		fmt.Println("All todos deleted")
	}
}

func DeleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/DeleteTodo/"))
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}

	for i, todo := range dto.Todos {
		if todo.ID == id {
			dto.Todos = append(dto.Todos[:i], dto.Todos[i+1:]...)
			fmt.Printf("Todo with ID %d deleted\n", id)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound) // delete todo by id
}

// func PutHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPut {

// 	}

// doesn`t need yet cuz of existing input`

// }

func main() {
	// db.TestDatabase() // doesn't need yet`
	http.HandleFunc("/DeleteTodo/", DeleteByIdHandler)
	http.HandleFunc("/DeleteTodo", DeleteHandler)
	http.HandleFunc("/GetTodo/", GetByIdHandler)
	http.HandleFunc("/GetTodo", GetHandler)
	http.HandleFunc("/todo", PostHandler)
	http.HandleFunc("/", Handler) // Register the handler for all paths
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the server
}
