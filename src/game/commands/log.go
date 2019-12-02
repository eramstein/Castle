package commands

import (
	m "github.com/castle/src/game/model"
)

func ClearLog(gs *m.State) {
	gs.Log.Logs = nil
	gs.Log.UserFeedback = false
}
