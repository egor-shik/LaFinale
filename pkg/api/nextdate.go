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
		return nextDay(now, date, parts)

    case "w":
    return nextWeek(now, date, parts)

    case "m":
    return nextMonth(now, date, parts)
    
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
    func nextDay(now, date time.Time, parts []string) (string, error) {
        if len(parts) != 2 {
            return "", errors.New("invalid d format")
        }
    
        days, err := strconv.Atoi(parts[1])
        if err != nil || days <= 0 || days > 400 {
            return "", errors.New("invalid days")
        }
        for {
            date = date.AddDate(0, 0, days)
            if date.After(now) {
                return date.Format(dateFormat), nil
            }
        }
    }
    
    func nextWeek(now, start time.Time, parts []string) (string, error) {
        if len(parts) != 2 {
            return "", errors.New("invalid w format")
        }

        days := parseInts(parts[1])
        if days == nil {
            return "", errors.New("invalid w format")
        }

        valid := make(map[int]bool)
        for _, d := range days {
            if d < 1 || d > 7 {
                return "", errors.New("invalid w format")
            }
            valid[d] = true
        }
        for i := 1; i <= 730; i++ {
            d := start.AddDate(0, 0, i)
    
            if d.After(now) {
                wd := int(d.Weekday())
                if wd == 0 {
                    wd = 7
                }
                if valid[wd] {
                    return d.Format(dateFormat), nil
                }
            }
        }
    
        return "", errors.New("no date found")
    }
    
        func nextMonth(now, start time.Time, parts []string) (string, error) {
        if len(parts) < 2 {
            return "", errors.New("invalid m format")
        }
    
        days := parseInts(parts[1])
        if days == nil {
            return "", errors.New("invalid m format")
        }
        negative := []int{}
        
        for _, d := range days {
            if d == 0 || d < -31 || d > 31 {
                return "", errors.New("invalid m format")
            }
            if d < 0 {
                negative = append(negative, d)
            }
        }
        if len(negative) == 2 {
            minusOne := false
            for _, d := range negative {
                if d == -1 {
                    minusOne = true
                    break
                }
            }
            if !minusOne {
                return "", errors.New("invalid m format")
            }
        } else if len(negative) > 2 {
            return "", errors.New("invalid m format")
        }
    
        var months []int
        if len(parts) >= 3 {
            months = parseInts(parts[2])
            if months == nil {
                return "", errors.New("invalid m format")
            }
    
            for _, m := range months {
                if m < 1 || m > 12 {
                    return "", errors.New("invalid m format")
                }
            }
        }
        for i := 0; i <= 24; i++ {
            cur := start.AddDate(0, i, 0)
            if len(months) > 0 {
                monthOk := false
                for _, m := range months {
                    if int(cur.Month()) == m {
                        monthOk = true
                        break
                    }
                }
                if !monthOk {
                    continue
                }
            }
    
            lastDay := lastDayOfMonth(cur)
            var maybe []time.Time
    
            for _, d := range days {
                var day int
                if d > 0 {
                    day = d
                } else {
                    day = lastDay + d + 1
                }
                
                if day < 1 || day > lastDay {
                    continue
                }
                
                cand := time.Date(
                    cur.Year(),
                    cur.Month(),
                    day,
                    0, 0, 0, 0,
                    time.UTC,
                )
                
                if cand.After(now) {
                    maybe = append(maybe, cand)
                }
            }
    
            if len(maybe) > 0 {
                nrb := maybe[0]
                for _, c := range maybe {
                    if c.Before(nrb) {
                        nrb = c
                    }
                }
                return nrb.Format(dateFormat), nil
            }
        }
    
        return "", errors.New("no date found")
    }
    
    func parseInts(s string) []int {
        parts := strings.Split(s, ",")
        var res []int
    
        for _, p := range parts {
            n, err := strconv.Atoi(strings.TrimSpace(p))
            if err != nil {
                return nil 
            }
            res = append(res, n)
        }
    
        return res
    }

    func lastDayOfMonth(t time.Time) int {
        return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
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