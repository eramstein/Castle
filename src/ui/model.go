package ui

import (
	m "github.com/castle/src/game/model"
)

type State struct {
	Camera  *Camera
	Texts   []*UiElement
	Buttons []*UiElement
}

type Camera struct {
	Width  int
	Height int
	Pos    *m.Pos
}

type Action struct {
	Name string
	Data int
}
