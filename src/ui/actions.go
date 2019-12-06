package ui

import (
	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
)

func handleAction(gs *m.State, ui *State, action Action) {
	switch action.Name {
	case "toggleInfoDetails":
		toggleInfoDetails(ui, action.EntityType, action.Entity, 0, 0)
	case "eat":
		cmd.Eat(gs, action.Data, action.Entity)
	case "pickup":
		cmd.Pickup(gs, action.Data, action.EntityType, action.Entity)
	}
}
