package main

import (
	blt "bearlibterminal"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"

	g "github.com/castle/src/game"
	"github.com/castle/src/ui"
)

type FullState struct {
	Game *g.State
	UI   *ui.State
}

var State = &FullState{}

const (
	WindowSizeX   = ui.CameraDefaultWidth*TileSizeX + ui.InfoPanelDefaultWidth*TextSizeX
	WindowSizeY   = ui.CameraDefaultHeight * TileSizeY
	Title         = "Château Ramstein"
	Font          = "UbuntuMono.ttf"
	CellPixelSize = 10
	TileSizeX     = 2
	TileSizeY     = 2
	TextSizeX     = 2
	TextSizeY     = 3
)

func init() {

	State = loadState()

	if State.Game == nil {
		State.Game = g.InitGame()
	}

	State.UI = ui.InitUI(State.Game)

	blt.Open()

	size := "size=" + strconv.Itoa(WindowSizeX) + "x" + strconv.Itoa(WindowSizeY)
	title := "title='" + Title + "'"
	cellsize := "cellsize=" + strconv.Itoa(CellPixelSize) + "x" + strconv.Itoa(CellPixelSize)
	window := "window: " + size + "," + title + "," + cellsize

	input := "input.filter = [keyboard, mouse]"

	blt.Set(window + "; " + input)
	blt.Set("text font: " + Font + ",size=16,spacing=1x2;")
	blt.Set("tile font: " + Font + ",size=18;")

}

func main() {

	ui.RenderAll(State.Game, State.UI)

	for {

		blt.Refresh()
		key := blt.Read()

		if key != blt.TK_CLOSE {
			if key == blt.TK_S {
				saveState()
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

func saveState() {
	runtime.LockOSThread()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(State)
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
