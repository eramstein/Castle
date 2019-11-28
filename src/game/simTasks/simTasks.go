package simTasks

import (
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func GetInitialSimTasks() []m.SimTaskSchedule {
	tasks := make([]m.SimTaskSchedule, 0)
	tasks = append(tasks, m.SimTaskSchedule{
		EndTime: 3600,
		Type:    m.TaskTypeUpdateNeeds,
	})
	tasks = append(tasks, m.SimTaskSchedule{
		EndTime: 4 * 3600,
		Type:    m.TaskTypeUpdateWeather,
	})
	return tasks
}

func AddPlayerTask(gs *m.State, task m.CharacterTask, estimatedDuration int) {
	gs.Player.Task = task
	gs.SimTasks = append(gs.SimTasks, m.SimTaskSchedule{
		EndTime: gs.Time + estimatedDuration,
		Type:    task.Type,
		Agent:   c.PlayerAgentID,
	})
}

func RunSimulation(gs *m.State, duration int) {
	taskIndex := getNextTask(gs.SimTasks)
	if len(gs.SimTasks) == 0 || gs.SimTasks[taskIndex].EndTime > gs.Time+duration {
		gs.Time += duration
		return
	}
	// execute first task && apply side effects on other tasks
	timeSpent := executeTask(gs, gs.SimTasks[taskIndex])
	// update time
	gs.Time += timeSpent
	remainingDuration := duration - timeSpent
	// remove task
	gs.SimTasks[taskIndex] = gs.SimTasks[len(gs.SimTasks)-1]
	gs.SimTasks = gs.SimTasks[:len(gs.SimTasks)-1]
	// loop to next task with remaining time
	if remainingDuration > 0 {
		RunSimulation(gs, remainingDuration)
	}
}

func getNextTask(tasks []m.SimTaskSchedule) (index int) {
	minTime := tasks[0].EndTime
	idx := 0
	for i, v := range tasks {
		if v.EndTime < minTime {
			minTime = v.EndTime
			idx = i
		}
	}
	return idx
}
