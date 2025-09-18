package dto

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
type TodoPost struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

var Todos []TodoPost
var IDTodo = 1
