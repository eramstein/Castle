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
	m "github.com/castle/src/game/model"
	"github.com/castle/src/ui"
)

type FullState struct {
	game *m.State
	ui   *ui.State
}

var State = &FullState{}

const (
	WindowSizeX   = ui.CameraDefaultWidth
	WindowSizeY   = ui.CameraDefaultHeight
	Title         = "Château Ramstein"
	Font          = "UbuntuMono.ttf"
	FontSize      = 24
	CellPixelSize = FontSize
)

func init() {

	State = loadState()

	if State.game == nil {
		State.game = g.InitGame()
	}

	State.ui = ui.InitUI(State.game)

	blt.Open()

	size := "size=" + strconv.Itoa(WindowSizeX) + "x" + strconv.Itoa(WindowSizeY)
	title := "title='" + Title + "'"
	cellsize := "cellsize=" + strconv.Itoa(CellPixelSize) + "x" + strconv.Itoa(CellPixelSize)
	window := "window: " + size + "," + title + "," + cellsize

	fontSize := "size=" + strconv.Itoa(FontSize)
	font := "font: " + Font + ", " + fontSize

	fmt.Println(State)

	blt.Set(window + "; " + font)

}

func main() {

	ui.RenderAll(State.game, State.ui)

	for {

		blt.Refresh()
		key := blt.Read()

		if key != blt.TK_CLOSE {
			handleInput(key)
		} else {
			break
		}

		ui.RenderAll(State.game, State.ui)

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

func handleInput(key int) {
	switch key {
	case blt.TK_RIGHT:
	case blt.TK_LEFT:
	case blt.TK_UP:
	case blt.TK_DOWN:
	case blt.TK_S:
		saveState()
	case blt.TK_L:
		loadState()
	}

}
