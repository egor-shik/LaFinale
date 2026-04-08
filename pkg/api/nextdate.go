package api 

import (
    "errors"
    "strings"
    "strconv"
    "time"
    "net/http"

)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
    dateDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
    nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
    
    return dateDay.After(nowDay)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

    if repeat == "" {
        return "", errors.New("repeat is empty")
    }

    date, err := time.Parse(dateFormat, dstart)
    if err != nil {
        return "", err
    }

    parts := strings.Split(repeat, " ")

    switch parts[0] {

    case "d":
        if len(parts) != 2 {
            return "", errors.New("invalid days")
        }
        days, err := strconv.Atoi(parts[1])
        if err != nil || days <= 0 || days > 400 {
            return "", errors.New("invalid days")
        }
        for {
                date = date.AddDate(0, 0, days)
                if afterNow(date, now) {
                    break
                }
        }
    case "y":
        for {
            date = date.AddDate(1, 0, 0)
            if afterNow(date, now) {
                break
            }
        }
    default:
        return "", errors.New("unsupported formait of repeat")
    }
    return date.Format(dateFormat), nil
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
    nowStr := r.FormValue("now")
    dateStr := r.FormValue("date")
    repeat := r.FormValue("repeat")

    var now time.Time
    var err error

    if nowStr == "" {
        now = time.Now()
    } else {
        now, err = time.Parse(dateFormat, nowStr)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
    }

    result, err := NextDate(now, dateStr, repeat)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.Write([]byte(result))
}