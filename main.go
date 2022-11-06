package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

var (
	todos []Todo
)

func main() {
	todos = append(todos, Todo{
		Title:  "todo 1",
		IsDone: false,
	})
	todos = append(todos, Todo{
		Title:  "todo 2",
		IsDone: false,
	})
	todos = append(todos, Todo{
		Title:  "todo 3",
		IsDone: false,
	})

	http.HandleFunc("/", getTodos)
	http.ListenAndServe(":8080", nil)

}

func getTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	case "POST":
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}
