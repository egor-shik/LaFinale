package api

import (
	"net/http"

	"LaFinale/pkg/db"
)

const defaultTasksLimit = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(defaultTasksLimit)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
