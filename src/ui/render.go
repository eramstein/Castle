package ui

import (
	blt "bearlibterminal"
	"fmt"
	"strconv"

	m "github.com/castle/src/game/model"
)

var i = 0

func RenderAll(gs *m.State, ui *State) {
	i++
	fmt.Println("render all " + strconv.Itoa(i))
	blt.Clear()
	renderMap(ui.Camera, gs.World)
	renderPlayer(ui.Camera, gs.Player)
	renderInfoPanel(ui.Texts, ui.Buttons)
}
