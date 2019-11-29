package model

type Log struct {
	SimTimeLeft int
	Logs        []LogForUI
}

type LogForUI struct {
	LogType int
	Text    string
}

const (
	LogTypeImpossibleCommand = 0
)
