package ui

import (
	"fmt"

	m "github.com/castle/src/game/model"
)

func RenderAll(gs *m.State, ui *State) {
	fmt.Println("render all", ui.Camera.Pos.X)
	renderMap(ui.Camera, gs.World)
	renderPlayer(ui.Camera, gs.Player)
}
