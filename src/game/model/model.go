package model

type State struct {
	Player *Player
	World  *World
	Time   int
}

type Player struct {
	Pos *Pos
}
