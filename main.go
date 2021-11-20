package main

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/henrikac/learn-go-swagger/todos"
)

func main() {
	storage := todos.NewMemoryStorage()
	th := todos.Handler{
		DB: storage,
	}

	http.HandleFunc("/api/todos", th.GetCreate)
	http.HandleFunc("/api/todos/", th.GetUpdateDelete)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	sh := middleware.Redoc(opts, nil)

	http.Handle("/docs", sh)
	http.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
