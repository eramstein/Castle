package ui

import (
	g "github.com/castle/src/game"
)

type State struct {
	Camera  *Camera
	Texts   []*UiElement
	Buttons []*UiElement
}

type Camera struct {
	Width  int
	Height int
	Pos    *g.Pos
}

type Action struct {
	name string
	data int
}
