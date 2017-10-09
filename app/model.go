package app

import (
	"database/sql"
	_ "errors"
)

type todo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueTo       int64  `json:"dueTo"`
}

func (todo *todo) getTodo(db *sql.DB) error {
	return db.QueryRow("SELECT name, description, due_to FROM todo WHERE id=$1",
		todo.ID).Scan(&todo.Name, &todo.Description, &todo.DueTo)
}

func (todo *todo) updateTodo(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE todo SET name=$1, description=$2, due_to=$3 WHERE id=$4",
			todo.Name, todo.Description, todo.DueTo, todo.ID)

	return err
}

func (todo *todo) deleteTodo(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM todo WHERE id=$1", todo.ID)

	return err
}

func (todo *todo) createTodo(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO todo(name, description, due_to) VALUES($1, $2, $3) RETURNING id",
		todo.Name, todo.Description, todo.DueTo).Scan(&todo.ID)

	if err != nil {
		return err
	}

	return nil
}

func getTodo(db *sql.DB, start, count int) ([]todo, error) {
	rows, err := db.Query(
		"SELECT id, name, description, due_to FROM todo LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todoArr := []todo{}

	for rows.Next() {
		var todo todo
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.DueTo); err != nil {
			return nil, err
		}
		todoArr = append(todoArr, todo)
	}

	return todoArr, nil
}
