package commands

import (
	char "github.com/castle/src/game/character"
	m "github.com/castle/src/game/model"
	simTasks "github.com/castle/src/game/simTasks"
	"github.com/davecgh/go-spew/spew"
)

func MovePlayer(gs *m.State, x, y, z int) {
	// check for validity
	// estimate task duration
	var pos = &gs.Player.Pos
	tile := gs.World.Regions[pos.Region].Tiles[pos.Z+z][pos.Y+y][pos.X+x]
	dur := char.GetTimeMoveToTile(gs.Player.Atts, gs.Player.Physical, tile)
	// create task & add to sim schedule and player data
	task := m.CharacterTask{
		LastUpdated: gs.Time,
		Completion:  0,
		Type:        m.TaskTypeMovement,
		Pos: m.Pos{
			Region: pos.Region,
			X:      pos.X + x,
			Y:      pos.Y + y,
			Z:      pos.Z + z,
		},
	}
	simTasks.AddPlayerTask(gs, task, dur)
	// run simulation during task time or until interruption
	simTasks.RunSimulation(gs, dur)
	spew.Dump("SCHEDULE", gs.SimTasks)
}
