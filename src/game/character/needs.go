package character

import (
	"fmt"

	c "github.com/castle/src/game/config"
	"github.com/castle/src/game/item"
	log "github.com/castle/src/game/log"
	m "github.com/castle/src/game/model"
	u "github.com/castle/src/game/utils"
)

const (
	HUNGER_NEED_RATE   = 10 // number of hours for hunger to go from 0 to 100
	THIRST_NEED_RATE   = 6  // number of hours for thirst to go from 0 to 100
	HUNGER_DAMAGE_RATE = 1  // damage done to nutrition if hunger >= 100
	THIRST_DAMAGE_RATE = 2  // damage done to hydratation if thirst >= 100
	HUNGER_REPAIR_RATE = 1  // repair done to nutrition if hunger < 100
	THIRST_REPAIR_RATE = 1  // repair done to hydratation if thirst < 100
)

func GetTimeToEat() (seconds int) {
	return 300
}

func GetTimeToDrink() (seconds int) {
	return 30
}

func UpdateNeeds(gs *m.State) {
	fmt.Println("UpdateNeeds")
	for _, char := range gs.Characters {
		updateNeedsState(gs, &char.Needs, &char.Physical, char.NeedsProfile)
	}
	gs.Log.UserFeedback = true
}

func updateNeedsState(gs *m.State, needsState *m.NeedsState, physical *m.PhysicalState, profile m.NeedsProfile) {
	// HUNGER
	var hungerRate = 100 * (float64(c.UpdateNeedsFrequency) / (3600 * HUNGER_NEED_RATE))
	needsState.Hunger += int(hungerRate * float64(profile.Hunger) / 100)
	if needsState.Hunger > 100 {
		physical.Nutrition.Total -= HUNGER_DAMAGE_RATE
		if physical.Nutrition.Total <= 0 {
			physical.Alive = false
		}
		needsState.Hunger = 100
		log.AddLog(gs, m.LogForUI{
			LogType: m.LogTypeAlert,
			Text:    "faim",
		})
	} else {
		if physical.Nutrition.Total < 100 {
			physical.Nutrition.Total += HUNGER_REPAIR_RATE
		}
	}
	// THIRST
	var thirstRate = 100 * (float64(c.UpdateNeedsFrequency) / (3600 * THIRST_NEED_RATE))
	needsState.Thirst += int(thirstRate * float64(profile.Thirst) / 100)
	if needsState.Thirst > 100 {
		physical.Hydratation -= THIRST_DAMAGE_RATE
		if physical.Hydratation <= 0 {
			physical.Alive = false
		}
		needsState.Thirst = 100
		log.AddLog(gs, m.LogForUI{
			LogType: m.LogTypeAlert,
			Text:    "soif",
		})
	} else {
		if physical.Hydratation < 100 {
			physical.Hydratation += THIRST_REPAIR_RATE
		}
	}
}

func Eat(gs *m.State, agentID int, quantity int, tile *m.Tile, where int, index int) {
	// update food slice
	item.UpdateStackQuantity(gs, -quantity, tile, agentID, where, index, c.ItemTypeFood)
	// update entity needs
	gs.Characters[agentID].Needs.Hunger = u.Max(gs.Characters[agentID].Needs.Hunger-10, 0)
}

func Drink(gs *m.State, agentID int, quantity int, tile *m.Tile, where int, index int) {
	// case from tile directly (e.g. river)
	if where == c.WhereFloor {
		if tile.Surface != m.SurfaceWater && tile.Volume != m.VolumeWater {
			return
		}
	}

	// TODO: case from container
	//item.UpdateStackQuantity(gs, -quantity, tile, 0, where, index, c.ItemTypeFood)

	// update entity needs
	gs.Characters[agentID].Needs.Thirst = u.Max(gs.Characters[agentID].Needs.Thirst-10, 0)
}
