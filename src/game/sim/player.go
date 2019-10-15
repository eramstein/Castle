package sim

import (
	m "github.com/castle/src/game/model"
)

func InitPlayer() *m.Player {
	return &m.Player{
		Pos: &m.Pos{
			X:      15,
			Y:      15,
			Z:      0,
			Region: 0,
		},
	}
}
