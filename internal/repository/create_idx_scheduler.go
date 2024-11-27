package repository

func (r *Repository) CreateIDXScheduler() error {
	query := `CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);`
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
