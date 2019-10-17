package ui

import (
	g "github.com/castle/src/game"
)

const (
	InfoPanelDefaultWidth = 30
	CameraDefaultWidth    = 61
	CameraDefaultHeight   = 61
)

type State struct {
	Camera *Camera
}

type Camera struct {
	Width  int
	Height int
	Pos    *g.Pos
}
