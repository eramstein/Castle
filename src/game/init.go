package game

import (
	m "github.com/castle/src/game/model"
	"github.com/castle/src/game/sim"
)

func InitGame() *m.State {
	var state = &m.State{
		Player: nil,
		World:  nil,
	}
	state.Player = sim.InitPlayer()
	state.World = sim.InitWorld()
	return state
}
