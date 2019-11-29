package character

import (
	"fmt"

	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func UpdateNeeds(gs *m.State) {
	fmt.Println("UpdateNeeds")
	updateNeedsState(&gs.Player.Needs, gs.Player.NeedsProfile)
}

func updateNeedsState(needsState *m.NeedsState, profile m.NeedsProfile) {
	var hungerRate = int(100 * (float64(c.UpdateNeedsFrequency) / (3600 * 10))) // hungers goes from 0 to 100 in 10 hours
	needsState.Hunger += int(hungerRate * profile.Hunger / 100)
	if needsState.Hunger > 100 {
		needsState.Hunger = 100
	}
}
