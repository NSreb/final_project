package repository

func (r *Repository) CreateScheduler() error {
	query := `CREATE TABLE IF NOT EXISTS scheduler (
              id INTEGER PRIMARY KEY AUTOINCREMENT
             ,date TEXT NOT NULL DEFAULT ""      
             ,title TEXT NOT NULL DEFAULT "" 
             ,comment TEXT NOT NULL DEFAULT "" 
             ,repeat TEXT NOT NULL DEFAULT "" );`
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
