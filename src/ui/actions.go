package ui

import (
	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
)

func handleAction(gs *m.State, ui *State, action *Action) {
	switch action.Name {
	case "movePlayer":
		cmd.MovePlayer(gs, 1, 0, 0)
	case "toggleInfoDetails":
		toggleInfoDetails(ui, action.EntityType, action.Entity)
	}
}
