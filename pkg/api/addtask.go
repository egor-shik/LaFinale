package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"strings"
	"errors"

	"LaFinale/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "title is required"})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, map[string]string{"id": strconv.FormatInt(id, 10)})
}

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" { //Чтобы избежать ситуации, когда дату не передали
		task.Date = now.Format(dateFormat)
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return err
	}

	if !isValidRepeat(task.Repeat) {
		return errors.New("unsupported repeat format")
	}

	if strings.HasPrefix(task.Repeat, "d ") || task.Repeat == "y" {   //Проверка на подсчет даты 
		_, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if !t.After(now) {
		task.Date = now.Format(dateFormat)
	}

	return nil
}
