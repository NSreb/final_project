package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) UpdDateTask(id int, date string) error {

	query := "update scheduler set date = :date WHERE id = :id"
	result, err := r.db.Exec(query, sql.Named("id", id), sql.Named("date", date))
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при получении количества затронутых строк: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}

	return nil
}
