package repository

import (
	"database/sql"
	"fmt"
	"go_final_project/internal/models"
)

type Task struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (r *Repository) GetTask(id models.TaskId) (*models.Tasks, error) {
	query := "SELECT ID, DATE, TITLE, COMMENT, REPEAT FROM SCHEDULER WHERE ID = :ID"
	row := r.db.QueryRow(query, sql.Named("ID", id.Id))

	var data models.Tasks

	err := row.Scan(&data.Id, &data.Date, &data.Title, &data.Comment, &data.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(" Задача с указаным ID не найдена")
		}
		return nil, err
	}

	return &data, nil
}
