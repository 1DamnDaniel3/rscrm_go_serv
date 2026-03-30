package usecaseutils

import "time"

// to find date by weekday
func GetDatesByWeekday(start, end time.Time, weekday time.Weekday) []time.Time {
	var dates []time.Time

	// 1. Найти первый нужный weekday
	offset := (int(weekday) - int(start.Weekday()) + 7) % 7
	first := start.AddDate(0, 0, offset)

	// 2. шаг 7 дней
	for d := first; !d.After(end); d = d.AddDate(0, 0, 7) {
		dates = append(dates, d)
	}

	return dates
}
