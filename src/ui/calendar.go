package ui

import (
	"strconv"
)

func getTimeString(timestamp int) string {
	day := int(timestamp / 86400)
	secondsInDay := timestamp - day*86400
	hour := int(secondsInDay / 3600)
	secondsInHour := timestamp - day*86400 - hour*3600
	minute := int(secondsInHour / 60)
	second := timestamp - day*86400 - hour*3600 - minute*60
	hourBuffer := ""
	if hour < 10 {
		hourBuffer = "0"
	}
	minuteBuffer := ""
	if minute < 10 {
		minuteBuffer = "0"
	}
	secondBuffer := ""
	if second < 10 {
		secondBuffer = "0"
	}
	return "Jour " + strconv.Itoa(day+1) +
		" - " + hourBuffer + strconv.Itoa(hour) +
		":" + minuteBuffer + strconv.Itoa(minute) +
		":" + secondBuffer + strconv.Itoa(second)
}
