package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"LaFinale/pkg/db"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		getTaskHandler(w, r)

	case http.MethodPut:
		updateTaskHandler(w, r)

	case http.MethodPost:
		addTaskHandler(w, r)

	case http.MethodDelete:
		deleteTaskHandler(w, r)

	default:
		writeError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
	}
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	writeJSON(w, task)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		writeError(w, errors.New("empty body"), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeError(w, errors.New("id is not specified"), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeError(w, errors.New("title is not specified"), http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeError(w, err, http.StatusNotFound)
		return
	}

	writeJSON(w, map[string]interface{}{})
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, errors.New("id is not specified"), http.StatusBadRequest)
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeError(w, err, http.StatusNotFound)
		return
	}

	writeJSON(w, map[string]interface{}{})
}
