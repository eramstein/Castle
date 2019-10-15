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
	}
	return state
}
