package ui

import (
	blt "bearlibterminal"
	"fmt"

	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
)

func HandleInput(key int, gs *m.State, ui *State) {

	if key == blt.TK_ESCAPE {
		if ui.IntendedAction != 0 {
			ui.IntendedAction = 0
		}
		return
	}

	if key == blt.TK_SPACE && ui.BlockSim {
		ui.BlockSim = false
		cmd.InterruptSim(gs)
		return
	}

	if key == blt.TK_ENTER {
		blt.Color(blt.ColorFromName("white"))
		_, t := blt.ReadStr(1, 1, 100)
		handleConsoleInput(t, gs)
		return
	}

	if ui.Screen == ScreenMap {
		handleScreenMapInputs(key, gs, ui)
	}

}

func handleScreenMapInputs(key int, gs *m.State, ui *State) {
	//fmt.Println(key)

	for _, action := range BasicPlayerActions {
		if action.Key == key {
			if ui.IntendedAction == action.Type {
				ui.IntendedAction = 0
			} else {
				ui.IntendedAction = action.Type
			}
			return
		}
	}

	switch key {

	case blt.TK_RIGHT:
		cmd.MovePlayer(gs, 1, 0, 0)
	case blt.TK_LEFT:
		cmd.MovePlayer(gs, -1, 0, 0)
	case blt.TK_UP:
		cmd.MovePlayer(gs, 0, -1, 0)
	case blt.TK_DOWN:
		cmd.MovePlayer(gs, 0, 1, 0)

	case blt.TK_KP_1:
		cmd.MovePlayer(gs, -1, 1, 0)
	case blt.TK_KP_2:
		cmd.MovePlayer(gs, 0, 1, 0)
	case blt.TK_KP_3:
		cmd.MovePlayer(gs, 1, 1, 0)
	case blt.TK_KP_4:
		cmd.MovePlayer(gs, -1, 0, 0)
	case blt.TK_KP_6:
		cmd.MovePlayer(gs, 1, 0, 0)
	case blt.TK_KP_7:
		cmd.MovePlayer(gs, -1, -1, 0)
	case blt.TK_KP_8:
		cmd.MovePlayer(gs, 0, -1, 0)
	case blt.TK_KP_9:
		cmd.MovePlayer(gs, 1, -1, 0)

	case blt.TK_MOUSE_LEFT:
		clickedOnMap, x, y := tileAtCoordinates(ui.Camera, blt.State(blt.TK_MOUSE_X), blt.State(blt.TK_MOUSE_Y))
		fmt.Println(x, y)
		if clickedOnMap {
			setInfoDetails(ui, EntityTypeTile, 0, x, y)
		} else {
			handleButtonsClick(gs, ui, blt.State(blt.TK_MOUSE_X), blt.State(blt.TK_MOUSE_Y))
		}

	}
}
