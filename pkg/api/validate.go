package api

import (
	"strconv"
	"strings"
)

// Проверка правильности выполнения повторений
// d, w, y - день, неделя, год соответственно
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
		return err == nil && n > 0 && n <= 400

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
		return len(parts) == 1

	default:
		return false
	}
}
