package todos

type Todo struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}

type Storage interface {
	Get(int) (Todo, error)
	GetAll() []Todo
	Create(string) Todo
	Update(int, string) (Todo, error)
	Delete(int) (Todo, error)
}
