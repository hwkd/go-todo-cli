package main

import (
	"fmt"
	"os"

	"github.com/hwkd/go-todo-cli/pkg/todo"
)

func main() {
	action := os.Args[1]
	args := os.Args[2:]
	switch action {
	case "add":
		fmt.Println("Add a new todo")
		parseAddArgs(args)
	case "list":
		fmt.Println("List all todos")
	case "done":
		fmt.Println("Mark a todo as done")
	case "undone":
		fmt.Println("Mark a todo as undone")
	case "delete":
		fmt.Println("Delete a todo")
	default:
		fmt.Println("Invalid action")
	}

	todo := todo.NewTodo("Task 1", "Create a todo app")
	fmt.Println(todo.Title)
}

// parseAddArgs parses the arguments for the add action.
// Options:
//
//	-t title [required]
//	-d description [optional]
func parseAddArgs(args []string) {
}
