package player

import (
	m "github.com/castle/src/game/model"
)

func GetInitialPlayer() *m.Player {
	return &m.Player{
		Pos: &m.Pos{
			X:      15,
			Y:      15,
			Z:      0,
			Region: 0,
		},
	}
}
