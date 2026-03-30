package usecaseutils

import "time"

// to get minimal date of a&b
func MinTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

// to get max date of a&b
func MaxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}
