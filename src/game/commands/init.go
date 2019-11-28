package commands

import (
	m "github.com/castle/src/game/model"
	"github.com/castle/src/game/player"
	"github.com/castle/src/game/simTasks"
	"github.com/castle/src/game/world"
)

func InitGame() *m.State {
	var state = &m.State{
		Time: 0,
	}
	state.Player = player.GetInitialPlayer()
	state.World = world.GetInitialWorld()
	state.SimTasks = simTasks.GetInitialSimTasks()
	return state
}
