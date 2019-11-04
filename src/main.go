package main

import (
	blt "bearlibterminal"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"

	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
	"github.com/castle/src/ui"
)

type FullState struct {
	Game *m.State
	UI   *ui.State
}

var State = &FullState{}

func init() {

	State = loadState()

	if State.Game == nil {
		State.Game = cmd.InitGame()
	}

	State.UI = ui.InitUI(State.Game)

	blt.Open()
	blt.Set(ui.BltConfig)

}

func main() {

	ui.RenderAll(State.Game, State.UI)

	for {

		blt.Refresh()
		key := blt.Read()

		if key != blt.TK_CLOSE {
			if key == blt.TK_S {
				saveState(State)
			} else {
				ui.HandleInput(key, State.Game, State.UI)
			}
		} else {
			break
		}

		ui.RenderAll(State.Game, State.UI)

	}

	blt.Close()
}

func saveState(state *FullState) {
	runtime.LockOSThread()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(state)
	if err != nil {
		fmt.Println("PANIC ENCODE")
		panic(err)
	}

	err = ioutil.WriteFile("save.txt", buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("PANIC SAVE")
		panic(err)
	}
	runtime.UnlockOSThread()
}

func loadState() *FullState {
	runtime.LockOSThread()
	fmt.Println("LOAD START")
	b, err := ioutil.ReadFile("save.txt")
	if err != nil {
		fmt.Println("PANIC LOAD")
		return &FullState{}
	}

	buf := bytes.NewBuffer(b)

	dec := gob.NewDecoder(buf)

	var state FullState
	err = dec.Decode(&state)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Println("LOAD END")
	runtime.UnlockOSThread()
	return &state
}
