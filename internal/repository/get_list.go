package repository

type Data struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (r *Repository) GetList() ([]Data, error) {
	query := `select id, date, title, comment, repeat from scheduler order by date desc`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Data
	for rows.Next() {
		var data Data
		if err := rows.Scan(&data.Id, &data.Date, &data.Title, &data.Comment, &data.Repeat); err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
