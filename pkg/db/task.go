package db

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}
 
func AddTask(task *Task) (int64, error) { //Задача отправляется в базу. Возвращается id новой записи
	query := `
	INSERT INTO scheduler (date, title, comment, repeat)
	VALUES (?, ?, ?, ?)
	`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
func Tasks(limit int) ([]*Task, error) {
	if limit <= 0 {
		limit = 50
	}
		rows, err := DB.Query(`
			SELECT id, date, title, comment, repeat
			FROM scheduler
			ORDER BY date
			LIMIT ?
		`, limit)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
	
		tasks := make([]*Task, 0, limit)
	
		for rows.Next() {
			var t Task
			err := rows.Scan(
				&t.ID,
				&t.Date,
				&t.Title,
				&t.Comment,
				&t.Repeat,
			)
			if err != nil {
				return nil, err
			}
	
			tasks = append(tasks, &t)
		}
		if tasks == nil {
			tasks = []*Task{}
		}
	
		return tasks, nil
	}