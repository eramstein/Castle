package main

import (
	blt "bearlibterminal"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"runtime"
	"time"

	cmd "github.com/castle/src/game/commands"
	m "github.com/castle/src/game/model"
	"github.com/castle/src/ui"
)

type FullState struct {
	Game *m.State
	UI   *ui.State
}

var State = &FullState{}
var iLoops = 0

func init() {

	rand.Seed(time.Now().UTC().UnixNano())

	runtime.LockOSThread()
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
				iLoops = 0
				loopSim()
			}
		} else {
			break
		}

		ui.RenderAll(State.Game, State.UI)

	}

	blt.Close()
}

func loopSim() {
	iLoops++
	if iLoops > 1000 {
		fmt.Println("TEMP safety vs infinite sim loops")
		return
	}
	ui.BlockSimIfNeeded(State.Game, State.UI)
	if State.Game.Log.SimTimeLeft > 0 && State.UI.BlockSim == false {
		time.Sleep(500 * time.Millisecond)
		ui.RenderAll(State.Game, State.UI)
		blt.Refresh()
		cmd.ResumeSim(State.Game)
		loopSim()
	}
}

func saveState(state *FullState) {

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

}

func loadState() *FullState {

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

	return &state
}
