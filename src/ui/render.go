package ui

import (
	blt "bearlibterminal"

	m "github.com/castle/src/game/model"
)

var i = 0

func RenderAll(gs *m.State, ui *State) {
	i++
	//fmt.Println("render all " + strconv.Itoa(i))
	blt.Clear()

	setUIElements(gs, ui)

	if ui.Screen == ScreenMap {
		renderMap(ui.Camera, gs.World)
		renderPlayer(ui.Camera, gs.Player)
		renderInfoPanel(ui.Texts, ui.Buttons)
	}

	blt.BkColor(blt.ColorFromName("black"))
}

func setUIElements(gs *m.State, ui *State) {
	ui.Texts = nil
	ui.Buttons = nil
	setInfoPanel(ui, gs)
}
