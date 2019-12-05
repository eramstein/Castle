package simTasks

import (
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func GetInitialSimTasks() []m.SimTaskSchedule {
	tasks := make([]m.SimTaskSchedule, 0)
	tasks = append(tasks, m.SimTaskSchedule{
		EndTime: c.UpdateNeedsFrequency,
		Type:    m.TaskTypeUpdateNeeds,
		Agent:   c.WorldAgentID,
	})
	tasks = append(tasks, m.SimTaskSchedule{
		EndTime: c.UpdateWeatherFrequency,
		Type:    m.TaskTypeUpdateWeather,
		Agent:   c.WorldAgentID,
	})
	return tasks
}

func AddPlayerTask(gs *m.State, task m.CharacterTask, estimatedDuration int) {
	gs.Characters[0].Task = task
	gs.SimTasks = append(gs.SimTasks, m.SimTaskSchedule{
		EndTime: gs.Time + estimatedDuration,
		Type:    task.Type,
		Agent:   0,
	})
}

func AddNeedsTask(gs *m.State, time int) {
	gs.SimTasks = append(gs.SimTasks, m.SimTaskSchedule{
		EndTime: time,
		Type:    m.TaskTypeUpdateNeeds,
		Agent:   c.WorldAgentID,
	})
}

func RunSimulation(gs *m.State, duration int) {
	if len(gs.SimTasks) == 0 {
		gs.Time += duration
		gs.Log.SimTimeLeft = 0
		return
	}
	taskIndex := getNextTask(gs.SimTasks)
	if gs.SimTasks[taskIndex].EndTime > gs.Time+duration {
		gs.Time += duration
		gs.Log.SimTimeLeft = 0
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
	// if no UI feedback needed, loop to next task with remaining time
	if gs.Log.UserFeedback == false && remainingDuration > 0 {
		RunSimulation(gs, remainingDuration)
	} else {
		gs.Log.SimTimeLeft = remainingDuration
	}
	return
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
