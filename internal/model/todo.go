package model

import (
	"fmt"
	"sync"
)

// Todo represents a single task item.
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// TodoStore is an in-memory datastore for Todos, safe for concurrent use.
type TodoStore struct {
	sync.Mutex
	todos  map[int]Todo
	nextID int
}

// NewTodoStore creates and initializes a new TodoStore.
func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]Todo),
		nextID: 1,
	}
}

// Create adds a new Todo to the store.
func (ts *TodoStore) Create(task string) Todo {
	ts.Lock()
	defer ts.Unlock()

	todo := Todo{
		ID:        ts.nextID,
		Task:      task,
		Completed: false,
	}
	ts.todos[todo.ID] = todo
	ts.nextID++
	return todo
}

// Get retrieves a single Todo from the store by its ID.
func (ts *TodoStore) Get(id int) (Todo, bool) {
	ts.Lock()
	defer ts.Unlock()

	todo, found := ts.todos[id]
	return todo, found
}

// List retrieves all Todos from the store.
func (ts *TodoStore) List() []Todo {
	ts.Lock()
	defer ts.Unlock()

	todos := make([]Todo, 0, len(ts.todos))
	for _, todo := range ts.todos {
		todos = append(todos, todo)
	}
	return todos
}

// Update modifies an existing Todo in the store.
func (ts *TodoStore) Update(id int, task string, completed bool) (Todo, bool) {
	ts.Lock()
	defer ts.Unlock()

	todo, found := ts.todos[id]
	if !found {
		return Todo{}, false
	}

	todo.Task = task
	todo.Completed = completed
	ts.todos[id] = todo
	return todo, true
}

// Delete removes a Todo from the store by its ID.
func (ts *TodoStore) Delete(id int) error {
	ts.Lock()
	defer ts.Unlock()

	if _, found := ts.todos[id]; !found {
		return fmt.Errorf("todo with id %d not found", id)
	}
	delete(ts.todos, id)
	return nil
}
