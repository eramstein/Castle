package ui

import (
	blt "bearlibterminal"

	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
)

func HandleInput(key int, gs *m.State, ui *State) {

	if ui.Screen == ScreenMap {
		handleScreenMapInputs(key, gs, ui)
	}

}

func handleScreenMapInputs(key int, gs *m.State, ui *State) {
	switch key {

	case blt.TK_RIGHT:
		cmd.MovePlayer(gs, 1, 0, 0)
	case blt.TK_LEFT:
		cmd.MovePlayer(gs, -1, 0, 0)
	case blt.TK_UP:
		cmd.MovePlayer(gs, 0, -1, 0)
	case blt.TK_DOWN:
		cmd.MovePlayer(gs, 0, 1, 0)

	case blt.TK_MOUSE_LEFT:
		clickedOnMap, x, y := tileAtCoordinates(ui.Camera, blt.State(blt.TK_MOUSE_X), blt.State(blt.TK_MOUSE_Y))
		if clickedOnMap {
			setInfoDetails(ui, EntityTypeTile, 0, x, y)
		} else {
			handleButtonsClick(gs, ui, blt.State(blt.TK_MOUSE_X), blt.State(blt.TK_MOUSE_Y))
		}

	}
}
