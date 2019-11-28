package model

type State struct {
	Player     *Player
	World      *World
	Time       int
	SimTasks   []SimTaskSchedule // not in order
	Characters []*Character
}
