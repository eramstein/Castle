package ui

import (
	"fmt"
	"strconv"

	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
)

func handleConsoleInput(text string, gs *m.State) {
	if len(text) <= 1 {
		return
	}
	key := text[0:1]
	rest := text[1:]

	switch key {
	case "w":
		time, err := strconv.Atoi(rest)
		if err != nil {
			fmt.Println("wesh bouffon")
			return
		}
		cmd.Wait(gs, time)
	}
}
