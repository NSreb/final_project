package repository

import (
	"database/sql"
	"fmt"
	"go_final_project/internal/models"
	"strconv"
)

func (r *Repository) EditTask(task models.Tasks) error {
	idStr := task.Id

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	query := "update scheduler set date = :date, title = :title, comment =:comment, repeat = :repeat WHERE id = :id"
	result, err := r.db.Exec(query, sql.Named("id", id), sql.Named("date", task.Date), sql.Named("title", task.Title), sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat))
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при получении количества затронутых строк: %w", err)
	}

	// Проверка, был ли обновлен хотя бы один ряд
	if rowsAffected == 0 {
		return fmt.Errorf("задача с ID %d не найдена", task.Id)
	}

	return nil
}
