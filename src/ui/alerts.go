package ui

import (
	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
)

func BlockSimIfNeeded(gs *m.State, ui *State) {
	block := false
	for _, logVal := range gs.Log.Logs {
		if logVal.LogType == m.LogTypeAlert {
			block = true
		}
	}
	ui.BlockSim = block
}

func clearAlerts(gs *m.State, ui *State) {
	cmd.ClearLog(gs)
}
