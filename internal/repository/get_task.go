package repository

import (
	"database/sql"
	"go_final_project/internal/models"
)

type Task struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (r *Repository) GetTask(id models.TaskId) (*Task, error) {
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id"
	row := r.db.QueryRow(query, sql.Named("id", id.Id))

	var data Task
	err := row.Scan(&data.Id, &data.Date, &data.Title, &data.Comment, &data.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &data, nil
}
