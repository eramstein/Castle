package simTasks

import (
	"fmt"

	char "github.com/castle/src/game/character"
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func executeTask(gs *m.State, task m.SimTaskSchedule) (timeSpent int) {
	charTask := m.CharacterTask{}
	if task.Agent != c.WorldAgentID {
		if task.Agent == c.PlayerAgentID {
			charTask = gs.Player.Task
		} else {
			charTask = gs.Characters[task.Agent].Task
		}
	}
	fmt.Println("task.Type", task.Type)
	switch task.Type {
	case m.TaskTypeMovement:
		char.MoveCharacter(gs, task.Agent, charTask.Pos)
	}

	return task.EndTime - gs.Time
}
