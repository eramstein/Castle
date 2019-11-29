package model

type State struct {
	Player     *Player
	World      *World
	Time       int
	SimTasks   []SimTaskSchedule
	Characters []*Character
	Log        Log
}
