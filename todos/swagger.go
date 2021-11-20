// Package classification Todo API.
//
// The purpose of this application is to learn about go-swagger.
//
// Schemes: http
// BasePath: /api
// Version: 0.0.1
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package todos

// swagger:route GET /todos todos listTodos
// Returns a list of all todos
// Responses:
//   200: todosResponse

// swagger:route GET /todos/{id} todos getTodo
// Returns the todo with the given id
// Responses:
//   200: todoResponse
//   400: badRequestResponse
//   404: notFoundResponse

// swagger:route POST /todos todos createTodo
// Creates a new todo
// Responses:
//   201: todoResponse
//   400: badRequestResponse

// swagger:route PUT /todos/{id} todos updateTodo
// Updates the todo with the given id
// Responses:
//   200: todoResponse
//   400: badRequestResponse
//   404: notFoundResponse

// swagger:route DELETE /todos/{id} todos deleteTodo
// Deletes the todo with the given id
// Responses:
//   200: todoResponse
//   400: badRequestResponse
//   404: notFoundResponse

// A list of todos returned in the response
// swagger:response todosResponse
type todosResponse struct {
	// All todos in the database
	// in: body
	Body []Todo
}

// A single todo returned in the response
// swagger:response todoResponse
type todoResponse struct {
	// The newly created todo
	// in: body
	Body Todo
}

// Todo task
// swagger:parameters createTodo updateTodo
type todoTaskParameter struct {
	// in: body
	Body struct {
		// The task of the todo
		// required: true
		Task string `json:"task"`
	}
}

// Todo id
// swagger:parameters deleteTodo getTodo updateTodo
type todoIdParameter struct {
	// The id of the todo
	// in: path
	Id int `json:"id"`
}

// Bad request
// swagger:response badRequestResponse
type badRequestResponse struct {
	// The error response
	// in: body
	Body struct {
		// The status code
		StatusCode int `json:"statusCode"`
		// The error message
		Message string `json:"message"`
	}
}

// Not found
// swagger:response notFoundResponse
type notFoundResponse struct {
	// The error response
	// in: body
	Body struct {
		// The status code
		StatusCode int `json:"statusCode"`
		// The error message
		Message string `json:"message"`
	}
}
