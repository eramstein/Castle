package simTasks

import (
	"fmt"
	"strconv"

	char "github.com/castle/src/game/character"
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
	"github.com/castle/src/game/world"
)

func executeTask(gs *m.State, task m.SimTaskSchedule) (timeSpent int) {
	charTask := m.CharacterTask{}
	if task.Agent != c.WorldAgentID {
		charTask = gs.Characters[task.Agent].Task
	}

	switch task.Type {
	case m.TaskTypeUpdateNeeds:
		char.UpdateNeeds(gs)
		AddNeedsTask(gs, task.EndTime+c.UpdateNeedsFrequency)
	case m.TaskTypeUpdateWeather:
		// TODO
	case m.TaskTypeMovement:
		char.MoveCharacter(gs, task.Agent, charTask.Pos)
	case m.TaskTypeEat:
		tile := world.GetTileAtPos(gs, charTask.Pos)
		char.Eat(gs, task.Agent, 1, tile, charTask.Where, charTask.ItemIndex)
	case m.TaskTypeDrink:
		tile := world.GetTileAtPos(gs, charTask.Pos)
		char.Drink(gs, task.Agent, 1, tile, charTask.Where, charTask.ItemIndex)
	case m.TaskTypePickup:
		tile := world.GetTileAtPos(gs, charTask.Pos)
		char.Pickup(gs, task.Agent, 1, tile, task.Agent, charTask.Where, charTask.ItemIndex, charTask.ItemType)
	default:
		fmt.Println("Task not found: " + strconv.Itoa(task.Type))
	}

	return task.EndTime - gs.Time
}
