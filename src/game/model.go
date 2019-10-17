package game

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
