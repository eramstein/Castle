package game

func InitGame() *State {
	var state = &State{
		Player: nil,
		World:  nil,
	}
	state.Player = InitPlayer()
	state.World = InitWorld()
	return state
}
