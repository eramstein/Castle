package ui

import (
	blt "bearlibterminal"
	"fmt"

	g "github.com/castle/src/game"
)

func HandleInput(key int, gs *g.State, ui *State) {
	switch key {
	case blt.TK_RIGHT:
		g.CmdMovePlayer(gs, 1, 0, 0)
	case blt.TK_LEFT:
		g.CmdMovePlayer(gs, -1, 0, 0)
	case blt.TK_UP:
		g.CmdMovePlayer(gs, 0, -1, 0)
	case blt.TK_DOWN:
		g.CmdMovePlayer(gs, 0, 1, 0)

	case blt.TK_MOUSE_LEFT:
		fmt.Println(blt.State(blt.TK_MOUSE_X))
	}

}
