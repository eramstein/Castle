package ui

import (
	"log"
	"time"
)

func runningtime(s string) (string, time.Time) {
	log.Println("Start:	", s)
	return s, time.Now()
}

func track(s string, startTime time.Time) {
	log.Println("End:	", s, "took", time.Since(startTime))
}
