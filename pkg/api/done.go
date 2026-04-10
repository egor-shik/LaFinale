package api

import (
	"errors"
	"net/http"
	"time"

	"LaFinale/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, errors.New("id is not specified"), http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, err, http.StatusNotFound)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeError(w, err, http.StatusInternalServerError)
			return
		}
		writeJSON(w, map[string]interface{}{})
		return
	}

	now := time.Now()

	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	err = db.UpdateDate(next, id)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]interface{}{})
}
