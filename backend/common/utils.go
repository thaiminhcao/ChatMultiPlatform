package common

import "time"

func GetCurrentTime() int64 {
	currentTime := time.Now()
	day := currentTime.Day()
	month := int(currentTime.Month())
	year := currentTime.Year()
	dateInt := int64(day*1000000 + month*10000 + year)
	return dateInt
}
