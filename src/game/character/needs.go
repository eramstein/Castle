package character

import (
	"fmt"

	c "github.com/castle/src/game/config"
	log "github.com/castle/src/game/log"
	m "github.com/castle/src/game/model"
	u "github.com/castle/src/game/utils"
)

func UpdateNeeds(gs *m.State) {
	fmt.Println("UpdateNeeds")
	for _, char := range gs.Characters {
		updateNeedsState(gs, &char.Needs, &char.Physical, char.NeedsProfile)
	}
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

func Eat(gs *m.State, agentID int, quantity int, pos m.Pos, where int, index int) {
	switch where {
	case c.WhereFloor:
		slice := gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y].Items.Food
		slice[index].Quantity -= quantity
		if slice[index].Quantity <= 0 {
			gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y].Items.Food = nil
		}
	}
	// update entity needs
	gs.Characters[agentID].Needs.Hunger = u.Max(gs.Characters[agentID].Needs.Hunger-10, 0)
}
