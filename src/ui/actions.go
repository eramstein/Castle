package ui

import (
	g "github.com/castle/src/game"
	"github.com/davecgh/go-spew/spew"
)

func handleAction(gs *g.State, action *Action) {
	spew.Dump(action)
	switch action.name {
	case "movePlayer":
		g.CmdMovePlayer(gs, 1, 0, 0)
	}
}
