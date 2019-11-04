package ui

import (
	m "github.com/castle/src/game/model"
)

func InitUI(gameState *m.State) *State {
	var player = gameState.Player
	var state = &State{
		Camera: &Camera{
			Width:  CameraDefaultWidth,
			Height: CameraDefaultHeight,
			Pos:    player.Pos,
		},
		Texts: []*UiElement{
			&UiElement{
				X:          5,
				Y:          30,
				Text:       "YOYOYOYOO",
				Width:      20,
				Height:     3,
				Color:      ColorRed,
				Background: ColorWhite,
			},
		},
		Buttons: []*UiElement{
			&UiElement{
				X:          5,
				Y:          40,
				Text:       "ITSAME CLICKAME",
				Width:      20,
				Height:     3,
				Color:      ColorRed,
				Background: ColorWhite,
				OnLeftClick: &Action{
					Name: "movePlayer",
					Data: 1,
				},
			},
		},
	}
	return state
}
