package ui

import (
	m "github.com/castle/src/game/model"
)

type State struct {
	Screen        int
	Camera        *Camera
	Texts         []*UiElement
	Buttons       []*UiElement
	EntityDetails Entity
}

type Camera struct {
	Width  int
	Height int
	Pos    *m.Pos
}

type Action struct {
	Name       string
	Data       int
	EntityType int
	Entity     int
}

type Entity struct {
	ID   int
	Type int
}

const (
	ScreenMap = iota
)

const (
	EntityTypeRegion = 1
)
