package model

type Log struct {
	UserFeedback bool
	SimTimeLeft  int
	Logs         []LogForUI
}

type LogForUI struct {
	LogType int
	Text    string
}

const (
	LogTypeImpossibleCommand = 0
	LogTypeAlert             = 1
)
