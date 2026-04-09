package api

import (
	"strings"
	"strconv"

	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error) {
	writeJSON(w, map[string]string{
		"error": err.Error(),
	})
}
//Проверка правильности выполнения повторений
//d, w, y - день, месяц, год соответственно
func isValidRepeat(r string) bool {
	if r == "" {
		return true
	}

	parts := strings.Split(r, " ")

	switch parts[0] {
	case "d":
		if len(parts) != 2 {
			return false
		}
		n, err := strconv.Atoi(parts[1])
		return err == nil && n > 0 && n <= 400 //диапазон ограничен тк 400 - максимальный срок переноса задачи 

	case "w":
		if len(parts) != 2 {
			return false
		}
		days := strings.Split(parts[1], ",")
		for _, d := range days {
			n, err := strconv.Atoi(d)
			if err != nil || n < 1 || n > 7 {
				return false
			}
		}
		return true
case "y":
	return len(parts) == 1 && parts[0] == "y"

default:
	return false
}
}