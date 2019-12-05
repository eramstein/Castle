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
	state.Characters = append(state.Characters, player.GetInitialPlayerCharacter())
	state.World = world.GetInitialWorld()
	state.SimTasks = simTasks.GetInitialSimTasks()
	return state
}
