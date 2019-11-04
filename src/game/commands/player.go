package commands

import (
	m "github.com/castle/src/game/model"
	"github.com/davecgh/go-spew/spew"
)

func MovePlayer(gs *m.State, x, y, z int) {
	player := gs.Player
	player.Pos.X += x
	player.Pos.Y += y
	player.Pos.Z += z
	spew.Dump(player)
}
