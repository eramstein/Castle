package ui

import (
	blt "bearlibterminal"
	"fmt"
	"strconv"

	g "github.com/castle/src/game"
)

var i = 0

func RenderAll(gs *g.State, ui *State) {
	i++
	fmt.Println("render all " + strconv.Itoa(i))
	blt.Clear()
	renderMap(ui.Camera, gs.World)
	renderPlayer(ui.Camera, gs.Player)
	renderInfoPanel(ui.Texts, ui.Buttons)
}
