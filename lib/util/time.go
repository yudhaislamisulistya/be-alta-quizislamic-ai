package util

import "time"

func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
