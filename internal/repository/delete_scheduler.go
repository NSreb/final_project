package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) DeleteTask(id int) error {
	query := `DELETE FROM scheduler where id = :id`
	result, err := r.db.Exec(query, sql.Named("id", id))
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
