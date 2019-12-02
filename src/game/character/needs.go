package character

import (
	"fmt"

	c "github.com/castle/src/game/config"
	log "github.com/castle/src/game/log"
	m "github.com/castle/src/game/model"
)

func UpdateNeeds(gs *m.State) {
	fmt.Println("UpdateNeeds")
	updateNeedsState(gs, &gs.Player.Needs, &gs.Player.Physical, gs.Player.NeedsProfile)
	gs.Log.UserFeedback = true
}

func updateNeedsState(gs *m.State, needsState *m.NeedsState, physical *m.PhysicalState, profile m.NeedsProfile) {
	// HUNGER
	var hungerRate = int(100 * (float64(c.UpdateNeedsFrequency) / (3600 * 10))) // hungers goes from 0 to 100 in 10 hours
	needsState.Hunger += int(hungerRate * profile.Hunger / 100)
	if needsState.Hunger > 100 {
		physical.Nutrition.Total -= (needsState.Hunger - 100)
		if physical.Nutrition.Total <= 0 {
			physical.Alive = false
		}
		needsState.Hunger = 100
		log.AddLog(gs, m.LogForUI{
			LogType: m.LogTypeAlert,
			Text:    "faim",
		})
	}
}
