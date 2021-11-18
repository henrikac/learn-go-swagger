package todos

import "fmt"

var (
	pk int = 1
)

type MemoryStorage struct {
	todos map[int]Todo
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[int]Todo),
	}
}

func (ms MemoryStorage) Get(id int) (Todo, error) {
	if todo, found := ms.todos[id]; found {
		return todo, nil
	}
	return Todo{}, fmt.Errorf("unable to retrieve todo with id %d", id)
}

func (ms MemoryStorage) GetAll() []Todo {
	todos := []Todo{}
	for _, todo := range ms.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (ms MemoryStorage) Create(task string) Todo {
	todo := Todo{
		Id:   pk,
		Task: task,
	}
	ms.todos[pk] = todo
	pk++
	return todo
}

func (ms MemoryStorage) Update(id int, task string) (Todo, error) {
	if _, found := ms.todos[id]; !found {
		return Todo{}, fmt.Errorf("todo with id %d not found", id)
	}
	todo := ms.todos[id]
	todo.Task = task
	return todo, nil
}

func (ms MemoryStorage) Delete(id int) (Todo, error) {
	if _, found := ms.todos[id]; !found {
		return Todo{}, fmt.Errorf("todo with id %d not found", id)
	}
	todo := ms.todos[id]
	delete(ms.todos, id)
	return todo, nil
}
