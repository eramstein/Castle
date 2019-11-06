package commands

import (
	m "github.com/castle/src/game/model"
)

func MovePlayer(gs *m.State, x, y, z int) {
	player := gs.Player
	player.Pos.X += x
	player.Pos.Y += y
	player.Pos.Z += z
	gs.Time += 30
}
