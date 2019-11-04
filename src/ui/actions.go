package ui

import (
	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
	"github.com/davecgh/go-spew/spew"
)

func handleAction(gs *m.State, action *Action) {
	spew.Dump(action)
	switch action.Name {
	case "movePlayer":
		cmd.MovePlayer(gs, 1, 0, 0)
	}
}
