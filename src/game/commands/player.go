package commands

import (
	m "github.com/castle/src/game/model"
)

func MovePlayer(gs *m.State, x, y, z int) {
	var pos = &gs.Player.Pos
	pos.X += x
	pos.Y += y
	pos.Z += z
	gs.Time += 30
}
