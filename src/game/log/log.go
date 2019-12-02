package log

import (
	m "github.com/castle/src/game/model"
)

func AddLog(gs *m.State, log m.LogForUI) {
	gs.Log.Logs = append(gs.Log.Logs, log)
	gs.Log.UserFeedback = true
}
