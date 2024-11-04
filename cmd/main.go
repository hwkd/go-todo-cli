package main

import (
	"fmt"
	"os"

	"github.com/hwkd/go-todo-cli/pkg/parser"
)

func main() {
	action := os.Args[1]
	args := os.Args[2:]

	parser := parser.NewParser(args)
	parsedArgs, err := parser.ParseArgs()
	if err != nil {
		panic(err)
	}

	// TODO: Implement the logic for each action.
	switch action {
	case "add":
		for _, arg := range parsedArgs {
			switch arg[0] {
			case "t":
				fmt.Println("Title:", arg[1])
			case "d":
				fmt.Println("Description:", arg[1])
			}
		}
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
}
