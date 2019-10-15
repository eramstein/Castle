package model

type State struct {
	Player *Player
	World  *World
}

type Player struct {
	Pos *Pos
}

type Pos struct {
	Region int
	X      int
	Y      int
	Z      int
}

type World struct {
	Regions []*Region
}

type Region struct {
	Name        string
	Description string
	Tiles       [][][]*Tile
}

type Tile struct {
	Surface int
}

const (
	SurfaceGround = 1
	SurfaceRock   = 2
)

var SurfaceNames = map[int]string{
	1: "Ground",
	2: "Rock",
}
