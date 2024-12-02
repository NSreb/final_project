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
	query := "UPDATE SCHEDULER SET DATE = :DATE, TITLE = :TITLE, COMMENT =:COMMENT, REPEAT = :REPEAT WHERE ID = :ID"
	result, err := r.db.Exec(query, sql.Named("ID", id), sql.Named("DATE", task.Date), sql.Named("TITLE", task.Title), sql.Named("COMMENT", task.Comment), sql.Named("REPEAT", task.Repeat))
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
