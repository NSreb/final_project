package repository

func (r *Repository) CreateIDXScheduler() error {
	query := `CREATE INDEX IF NOT EXISTS IDX_DATE ON SCHEDULER(DATE);`
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
