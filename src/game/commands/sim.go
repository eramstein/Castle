package commands

import (
	m "github.com/castle/src/game/model"
	simTasks "github.com/castle/src/game/simTasks"
)

func ResumeSim(gs *m.State) {
	simTasks.RunSimulation(gs, gs.Log.SimTimeLeft)
}

func Wait(gs *m.State, minutes int) {
	simTasks.RunSimulation(gs, 60*minutes)
}

func InterruptSim(gs *m.State) {
	gs.Log.SimTimeLeft = 0
}
