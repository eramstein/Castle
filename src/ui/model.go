package ui

import (
	blt "bearlibterminal"

	m "github.com/castle/src/game/model"
)

type State struct {
	Screen         int
	Camera         Camera
	Texts          []UiElement
	Buttons        []UiElement
	EntityDetails  Entity
	BlockSim       bool
	IntendedAction int
}

type Camera struct {
	Width  int
	Height int
	Pos    *m.Pos
}

type Action struct {
	Name       string
	Data       interface{}
	Data2      interface{}
	EntityType int
	Entity     int
}

type Entity struct {
	ID    int
	Type  int
	Data1 int
	Data2 int
}

const (
	ScreenMap = iota
)

const (
	EntityTypeRegion = 1
	EntityTypeTile   = 2
)

const (
	ActionEat = iota + 1
	ActionUseInventory
	ActionPickup
	ActionDrink
)

type BasicPlayerAction struct {
	Type   int
	Name   string
	Handle string
	Key    int
}

var BasicPlayerActions = []BasicPlayerAction{
	{ActionEat, "Manger", "m", blt.TK_M},
	{ActionUseInventory, "Inventaire", "i", blt.TK_I},
	{ActionPickup, "Ramasser", "r", blt.TK_R},
	{ActionDrink, "Boire", "b", blt.TK_B},
}
