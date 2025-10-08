package utils

import (
	"time"
)

func GetCurrentDateTime() (string, error) {
	const (
		TimeFormat = "2006-01-02 15:04:05.000"
		TimeZone   = "Asia/Bangkok"
	)
	loc, err := time.LoadLocation(TimeZone)
	if err != nil {
		return "", err
	}
	return time.Now().In(loc).Format(TimeFormat), nil
}
