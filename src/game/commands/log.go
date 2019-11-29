package commands

import (
	m "github.com/castle/src/game/model"
)

func AddLog(gs *m.State, log m.LogForUI) {
	gs.Log.Logs = append(gs.Log.Logs, log)
}

func ClearLog(gs *m.State) {
	gs.Log.Logs = nil
}
