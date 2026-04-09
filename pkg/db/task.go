package db

import (
	"database/sql"
	"errors"
	"fmt"
)

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

func GetTask(id string) (*Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = ?
	`

	var tmp Task 		//для разнообразия переменных в разных функциях 

	err := DB.QueryRow(query, id).Scan(
		&tmp.ID,
		&tmp.Date,
		&tmp.Title,
		&tmp.Comment,
		&tmp.Repeat,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("Task is not found")
	}
	if err != nil {
		return nil, err
	}

	return &tmp, nil
}

func UpdateTask(task *Task) error {
	query := `
	UPDATE scheduler
	SET date = ?, title = ?, comment = ?, repeat = ?
	WHERE id = ?
	`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Task is not found")
	}

	return nil
}

func UpdateDate(next string, id string) error {
	query := `
	UPDATE scheduler
	SET date = ?
	WHERE id = ?
	`

	res, err := DB.Exec(query, next, id)
	if err != nil {
		return err
	}

	udcount, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if udcount == 0 {
		return fmt.Errorf("Task os not found")
	}

	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = ?`

	res, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	dtcount, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if dtcount == 0 {
		return fmt.Errorf("Task is not found")
	}

	return nil
}

func Tasks(limit int) ([]*Task, error) {
	if limit <= 0 {
		limit = 50   // лимит. чтобы уйти от странных значений
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
				return nil, fmt.Errorf("scan task: %w", err)
		}
	
			tasks = append(tasks, &t)
		}

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("iterate tasks: %w", err)
		}
		return tasks, nil
	}