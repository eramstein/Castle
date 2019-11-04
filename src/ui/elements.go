package ui

import (
	blt "bearlibterminal"
	"fmt"

	m "github.com/castle/src/game/model"
)

type UiElement struct {
	X            int
	Y            int
	Height       int
	Width        int
	Color        int
	Background   int
	Text         string
	OnLeftClick  *Action
	OnRightClick *Action
}

func (b UiElement) onClick(gs *m.State) {
	handleAction(gs, b.OnLeftClick)
}

func (e UiElement) draw() {
	xStart := InforPanelStart
	blt.Color(blt.ColorFromName(Colors[e.Color]))
	datext := "[font=text]" + e.Text + "[/font]"
	blt.Print(xStart+e.X, e.Y, datext)
}

func renderElements(elements []*UiElement) {
	for _, e := range elements {
		e.draw()
	}
}

func handleButtonsClick(gs *m.State, ui *State, x, y int) {
	infoCellX := x - InforPanelStart
	fmt.Println(infoCellX, y)
	for _, b := range ui.Buttons {
		if b.X <= infoCellX && b.X+b.Width >= infoCellX && b.Y <= y && b.Y+b.Height >= y {
			b.onClick(gs)
		}
	}
}
