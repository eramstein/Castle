package ui

import (
	m "github.com/castle/src/game/model"
)

func handleAction(gs *m.State, ui *State, action Action) {
	switch action.Name {
	case "toggleInfoDetails":
		toggleInfoDetails(ui, action.EntityType, action.Entity, 0, 0)
	}
}
