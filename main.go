package main

import (
	"log"
	"net/http"

	"github.com/henrikac/learn-go-swagger/todos"
)

func main() {
	storage := todos.NewMemoryStorage()
	th := todos.Handler{
		DB: storage,
	}

	http.HandleFunc("/api/todos", th.GetCreate)
	http.HandleFunc("/api/todos/", th.GetUpdateDelete)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
