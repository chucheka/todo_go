package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	// APP GENERIC ROUTES
	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)

	// APP TODO ROUTES
	mux.HandleFunc("GET /v1/todoes", app.fetchTodoes)
	mux.HandleFunc("POST /v1/todoes", app.createTodo)
	mux.HandleFunc("GET /v1/todoes/{id}", app.fetchTodo)
	mux.HandleFunc("PUT /v1/todoes/{id}", app.updateTodo)
	mux.HandleFunc("DELETE /v1/todoes/{id}", app.deleteTodo)

	// APP USER ROUTES
	mux.HandleFunc("POST /v1/users", app.healthcheckHandler)
	mux.HandleFunc("PUT /v1/users/activate", app.healthcheckHandler)
	mux.HandleFunc("PUT /v1/users/password", app.healthcheckHandler)
	mux.HandleFunc("POST /v1/tokens/authenticate", app.healthcheckHandler)

	return mux

}
