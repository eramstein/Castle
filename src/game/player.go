package game

import (
	"github.com/davecgh/go-spew/spew"
)

type Player struct {
	Pos *Pos
}

func InitPlayer() *Player {
	return &Player{
		Pos: &Pos{
			X:      15,
			Y:      15,
			Z:      0,
			Region: 0,
		},
	}
}

func CmdMovePlayer(gs *State, x, y, z int) {
	player := gs.Player
	player.Pos.X += x
	player.Pos.Y += y
	player.Pos.Z += z
	spew.Dump(player)
}
