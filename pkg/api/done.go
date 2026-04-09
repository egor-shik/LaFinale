package api

import (
	"errors"
	"net/http"
	"time"

	"LaFinale/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, errors.New("id is not specified"))
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, err)
		return
	}


	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, map[string]interface{}{})
		return
	}

	now := time.Now()

	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.UpdateDate(next, id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]interface{}{})
}