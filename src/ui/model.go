package ui

import (
	m "github.com/castle/src/game/model"
)

const (
	CameraDefaultWidth  = 61
	CameraDefaultHeight = 41
)

type State struct {
	Camera *Camera
}

type Camera struct {
	Width  int
	Height int
	Pos    *m.Pos
}
