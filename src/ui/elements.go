package ui

import (
	blt "bearlibterminal"

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

func (b UiElement) onClick(gs *m.State, ui *State) {
	handleAction(gs, ui, b.OnLeftClick)
}

func (e UiElement) draw() {
	color := ColorBlackish
	if e.Color != 0 {
		color = e.Color
	}
	blt.Color(blt.ColorFromName(Colors[color]))
	text := "[font=text]" + e.Text + "[/font]"
	blt.PrintExt(e.X, e.Y, e.Width, e.Height, blt.TK_ALIGN_MIDDLE, text)
}

func renderElements(elements []*UiElement) {
	for _, e := range elements {
		e.draw()
	}
}

func handleButtonsClick(gs *m.State, ui *State, x, y int) {
	for _, b := range ui.Buttons {
		if b.X <= x && b.X+b.Width >= x && b.Y <= y && b.Y+b.Height >= y {
			b.onClick(gs, ui)
		}
	}
}
