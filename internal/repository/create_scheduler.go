package repository

func (r *Repository) CreateScheduler() error {
	query := `CREATE TABLE IF NOT EXISTS SCHEDULER (
              ID INTEGER PRIMARY KEY AUTOINCREMENT
             ,DATE TEXT NOT NULL DEFAULT ""      
             ,TITLE TEXT NOT NULL DEFAULT "" 
             ,COMMENT TEXT NOT NULL DEFAULT "" 
             ,REPEAT TEXT NOT NULL DEFAULT "" );`
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
