package api

import "net/http"

func taskHandler(w http.ResponseWriter, r *http.Request) { // switch для добавления кейсов в будущем
	switch r.Method {

	case http.MethodPost:
		addTaskHandler(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}