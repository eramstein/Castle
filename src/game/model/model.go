package model

type State struct {
	Player *Player
	World  *World
}

type Pos struct {
	Region int
	X      int
	Y      int
	Z      int
}

type Player struct {
	Pos *Pos
}
