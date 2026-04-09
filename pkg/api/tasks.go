package api

import (
	"net/http"
	"LaFinale/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) { //отдача списка задач
	tasks, err := db.Tasks(50)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}