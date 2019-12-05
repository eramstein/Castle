package commands

import (
	char "github.com/castle/src/game/character"
	log "github.com/castle/src/game/log"
	m "github.com/castle/src/game/model"
	simTasks "github.com/castle/src/game/simTasks"
)

func MovePlayer(gs *m.State, x, y, z int) {
	var pos = &gs.Characters[0].Pos
	targetPos := m.Pos{Region: pos.Region, X: pos.X + x, Y: pos.Y + y, Z: pos.Z + z}
	// check for validity
	valid := char.IsMoveValid(gs, targetPos)
	if valid == false {
		log.AddLog(gs, m.LogForUI{LogType: m.LogTypeImpossibleCommand, Text: "Can't move here"})
		return
	}
	tile := gs.World.Regions[pos.Region].Tiles[pos.Z+z][pos.Y+y][pos.X+x]
	// estimate task duration
	dur := char.GetTimeMoveToTile(gs.Characters[0].Atts, gs.Characters[0].Physical, tile)
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
}

func Eat(gs *m.State, food m.Food, quantity int, pos m.Pos, where int) {
	char.Eat(gs, 0, food, quantity, pos, where)
}
