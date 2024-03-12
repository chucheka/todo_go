package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Todoes TodoModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Todoes: TodoModel{DB: db},
	}
}
