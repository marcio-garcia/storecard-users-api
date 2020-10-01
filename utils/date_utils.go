package utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

// GetNow returns the current date and time in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns the formatted current date and time in UTC
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat returns the formatted date for DB storage
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
