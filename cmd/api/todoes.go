package main

import (
	"errors"
	"fmt"
	"github.com/chucheka/todo/internal/data"
	"github.com/chucheka/todo/internal/validator"
	"net/http"
)

func (app *application) fetchTodoes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "path: GET /v1/todoes")
}

func (app *application) fetchTodo(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	todo, err := app.models.Todoes.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, todo, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createTodo(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Text string   `json:"text"`
		Tag  []string `json:"tag"`
	}

	//err := json.NewDecoder(r.Body).Decode(&input)

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	todo := &data.Todo{
		Text: input.Text,
		Tag:  input.Tag,
	}

	v := validator.New()

	if data.ValidateTodo(v, todo); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Todoes.Insert(todo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/todo/%d", todo.Id))

	err = app.writeJSON(w, http.StatusCreated, todo, headers)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateTodo(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	todo, err := app.models.Todoes.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Text      *string  `json:"text"`
		Tag       []string `json:"tag"`
		Completed *bool    `json:"completed"`
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Text != nil {
		todo.Text = *input.Text
	}

	if input.Completed != nil {
		todo.Completed = *input.Completed
	}

	if input.Tag != nil {
		todo.Tag = input.Tag // Note that we don't need to dereference a slice.
	}

	v := validator.New()
	if data.ValidateTodo(v, todo); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Todoes.Update(todo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, todo, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Todoes.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, "todo successfully, deleted", nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
