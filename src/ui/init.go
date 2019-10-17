package ui

import (
	g "github.com/castle/src/game"
)

func InitUI(gameState *g.State) *State {
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
