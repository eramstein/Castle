package utils

import (
	"log"
	"time"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Runningtime(s string) (string, time.Time) {
	log.Println("Start:	", s)
	return s, time.Now()
}

func Track(s string, startTime time.Time) {
	log.Println("End:	", s, "took", time.Since(startTime).Milliseconds(), " ms")
}
