package todo

import (
	"time"
)

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTodo(title, description string) *Todo {
	return &Todo{
		ID:          -1, // -1 means not saved
		Title:       title,
		Description: description,
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (todo *Todo) Done() {
	todo.IsDone = true
}

func (todo *Todo) Undone() {
	todo.IsDone = false
}

type TodoList struct {
	Todos  []Todo
	Length int
}

func NewTodoList() *TodoList {
	return &TodoList{}
}

func (todoList *TodoList) List() []Todo {
	return todoList.Todos
}

func (todoList *TodoList) Get(id int) *Todo {
	for i := range todoList.Todos {
		todo := &todoList.Todos[i]
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

func (todoList *TodoList) Add(todo *Todo) {
	todoList.Todos = append(todoList.Todos, *todo)
	todoList.Length = len(todoList.Todos)
}

func (todoList *TodoList) Update(todo *Todo) {
	for i, _todo := range todoList.Todos {
		if _todo.ID == todo.ID {
			todoList.Todos[i] = *todo
			break
		}
	}
}

func (todoList *TodoList) Delete(id int) {
	for i := range todoList.Todos {
		todo := &todoList.Todos[i]
		if todo.ID == id {
			todoList.Todos = append(todoList.Todos[:i], todoList.Todos[i+1:]...)
			break
		}
	}
}
