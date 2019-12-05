package ui

import (
	blt "bearlibterminal"

	m "github.com/castle/src/game/model"
)

func RenderAll(gs *m.State, ui *State) {

	blt.Clear()

	setUIElements(gs, ui)

	if ui.Screen == ScreenMap {
		renderMap(ui.Camera, gs.World)
		renderPlayer(ui.Camera, gs.Characters[0])
		renderInfoPanel(ui.Texts, ui.Buttons)
	}

	clearAlerts(gs, ui)

	blt.BkColor(blt.ColorFromName("black"))
}

func setUIElements(gs *m.State, ui *State) {
	ui.Texts = nil
	ui.Buttons = nil
	setInfoPanel(ui, gs)
}
