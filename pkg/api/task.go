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
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
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
	writeJSON(w, task)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if r.Body == nil {
		writeError(w, errors.New("empty body"))
		return
	}

	if err != nil {
		writeError(w, err)
		return
	}

	if task.ID == "" {
		writeError(w, errors.New("id is not specified"))
		return
	}

	if task.Title == "" {
		writeError(w, errors.New("Title is not specified"))  //*Была ошибка, речь о title, разумеется
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]interface{}{})
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, errors.New("id is not specified"))
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, map[string]interface{}{})
}