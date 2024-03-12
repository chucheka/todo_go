package data

import (
	"database/sql"
	"errors"
	"github.com/chucheka/todo/internal/validator"
	"github.com/lib/pq"
	"time"
)

type Todo struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	Tag       []string  `json:"tag,omitempty"`
	Completed bool      `json:"completed"`
	Version   int32     `json:"-"`
}

type TodoModel struct {
	DB *sql.DB
}

func (t TodoModel) Insert(todo *Todo) error {

	query := `
			 INSERT INTO todoes (text,tag) 
			 VALUES ($1, $2)
			 RETURNING id, created_at, version`

	args := []interface{}{
		todo.Text,
		pq.Array(todo.Tag)}

	return t.DB.QueryRow(query, args...).Scan(&todo.Id, &todo.CreatedAt, &todo.Version)
}

func (t TodoModel) Get(id int64) (*Todo, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
 SELECT id, created_at, text,completed, tag, version
 FROM todoes WHERE id = $1`

	var todo Todo

	err := t.DB.QueryRow(query, id).Scan(
		&todo.Id,
		&todo.CreatedAt,
		&todo.Text,
		&todo.Completed,
		pq.Array(&todo.Tag),
		&todo.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &todo, nil
}

func (t TodoModel) Update(todo *Todo) error {

	query := `
				 UPDATE todoes 
				 SET text = $1, tag = $2, completed = $3, version = version + 1
				 WHERE id = $4
				 RETURNING version
				 `

	args := []interface{}{
		todo.Text,
		pq.Array(todo.Tag),
		todo.Completed,
		todo.Id,
	}

	return t.DB.QueryRow(query, args...).Scan(&todo.Version)
}

func (t TodoModel) Delete(id int64) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
 DELETE FROM todoes
 WHERE id = $1`

	result, err := t.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func ValidateTodo(v *validator.Validator, todo *Todo) {
	v.Check(todo.Text != "", "text", "must be provided")
	v.Check(len(todo.Text) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(todo.Tag != nil, "tag", "must be provided")
	v.Check(len(todo.Tag) >= 1, "tag", "must contain at least 1 tag")
	v.Check(len(todo.Tag) <= 5, "tag", "must not contain more than 5 tags")
	v.Check(validator.Unique(todo.Tag), "tag", "must not contain duplicate values")
}
