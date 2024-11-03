package todo

import (
	"fmt"
	"testing"
)

func TestNewTodo(t *testing.T) {
	todo := NewTodo("Task 1", "Created for test")
	// Unsaved todo ID should be -1.
	if todo.ID != -1 {
		t.Errorf("Expected -1, got %d", todo.ID)
	}
	// Title and Description should be equal to argument.
	if todo.Title != "Task 1" {
		t.Errorf("Expected Test, got %s", todo.Title)
	}
	if todo.Description != "Created for test" {
		t.Errorf("Expected Test, got %s", todo.Description)
	}
	// IsDone should default to false.
	if todo.IsDone != false {
		t.Errorf("Expected false, got %t", todo.IsDone)
	}
}

func TestNewTodoList(t *testing.T) {
	todoList := NewTodoList()
	expected_len := 0
	if len(todoList.Todos) != expected_len {
		t.Errorf("Expected %d, got %d", expected_len, len(todoList.Todos))
	}
}

func TestTodoListList(t *testing.T) {
	todoList := NewTodoList()
	expected_len := 5
	for i := 0; i < expected_len; i++ {
		todo := NewTodo(fmt.Sprintf("Task %d", i), "Created for test")
		todoList.Add(todo)
	}
	todos := todoList.List()
	if len(todos) != expected_len {
		t.Errorf("Expected %d, got %d", expected_len, len(todos))
	}
	for i, todo := range todos {
		expectedTitle := fmt.Sprintf("Task %d", i)
		expectedDesc := "Created for test"
		if todo.Title != expectedTitle {
			t.Errorf("Expected %s, got %s", expectedTitle, todo.Title)
		}
		if todo.Description != expectedDesc {
			t.Errorf("Expected %s, got %s", expectedDesc, todo.Description)
		}
	}
}

func TestTodoListGet(t *testing.T) {
	todoList := NewTodoList()
	for i := 0; i < 3; i++ {
		todo := NewTodo(fmt.Sprintf("Task %d", i), "Created for test")
		todo.ID = i
		todoList.Add(todo)
	}
	todo := todoList.Get(1)
	if todo.ID != 1 {
		t.Errorf("Expected 1, got %d", todo.ID)
	}
	if todo.Title != "Task 1" {
		t.Errorf("Expected Task 1, got %s", todo.Title)
	}
}

func TestTodoListAdd(t *testing.T) {
	todoList := NewTodoList()
	todo := NewTodo("Task 1", "Created for test")
	todoList.Add(todo)
	// Ensure todo is added.
	if len(todoList.Todos) != 1 {
		t.Errorf("Expected 1, got %d", len(todoList.Todos))
	}
}

func TestTodoListUpdate(t *testing.T) {
	todoList := NewTodoList()
	todo := NewTodo("Task 1", "Created for test")
	todoList.Add(todo)
	todo.Title = "Task 2"
	// Ensure todo is copied by value, not reference.
	if todoList.Todos[0].Title == "Task 2" {
		t.Errorf("Expected Task 1, got %s", todoList.Todos[0].Title)
	}
	todoList.Update(todo)
	// todoList should be updated.
	if todoList.Todos[0].Title != "Task 2" {
		t.Errorf("Expected Task 2, got %s", todoList.Todos[0].Title)
	}
	todo.Done()
	todoList.Update(todo)
	if todoList.Todos[0].IsDone != true {
		t.Errorf("Expected true, got %t", todo.IsDone)
	}
}

func TestTodoDone(t *testing.T) {
	todo := NewTodo("Task 1", "Created for test")
	todo.Done()
	if todo.IsDone != true {
		t.Errorf("Expected true, got %t", todo.IsDone)
	}
	todo.Undone()
	if todo.IsDone != false {
		t.Errorf("Expected false, got %t", todo.IsDone)
	}
}

func TestTodoDelete(t *testing.T) {
	todoList := NewTodoList()
	origin_len := 5
	for i := 0; i < origin_len; i++ {
		todo := NewTodo(fmt.Sprintf("Task %d", i), fmt.Sprintf("Desc %d", i))
		todo.ID = i
		todoList.Add(todo)
	}

	id := 2
	todoList.Delete(id)
	new_len := origin_len - 1
	if len(todoList.Todos) != new_len {
		t.Errorf("Expected %d, got %d", new_len, len(todoList.Todos))
	}
	for _, todo := range todoList.Todos {
		if todo.ID == id {
			t.Errorf("ID %d should be deleted", id)
		}
	}
}
