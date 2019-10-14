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

	"github.com/castle/src/ui"
)

const (
	WindowSizeX   = 60
	WindowSizeY   = 40
	MapWidth      = 1000
	MapHeight     = 1000
	Title         = "Château Ramstein"
	Font          = "UbuntuMono.ttf"
	FontSize      = 24
	CellPixelSize = FontSize
)

var (
	player   *ui.GameEntity
	entities []*ui.GameEntity
	gameMap  *ui.GameMap
)

func init() {
	fmt.Println("INIT START")
	player = &ui.GameEntity{X: 1, Y: 1, Layer: 1, Char: "@", Color: "white"}
	npc := &ui.GameEntity{X: 30, Y: 10, Layer: 0, Char: "N", Color: "red"}
	entities = append(entities, player, npc)

	gameMap = &ui.GameMap{Width: MapWidth, Height: MapHeight}
	gameMap.InitializeMap()

	blt.Open()

	size := "size=" + strconv.Itoa(WindowSizeX) + "x" + strconv.Itoa(WindowSizeY)
	title := "title='" + Title + "'"
	cellsize := "cellsize=" + strconv.Itoa(CellPixelSize) + "x" + strconv.Itoa(CellPixelSize)
	window := "window: " + size + "," + title + "," + cellsize

	fontSize := "size=" + strconv.Itoa(FontSize)
	font := "font: " + Font + ", " + fontSize

	blt.Set(window + "; " + font)
	blt.Clear()
	fmt.Println("INIT END")
}

func main() {
	// Main game loop
	renderAll()

	i := 0

	for {
		i++
		fmt.Println("FOR START " + strconv.Itoa(i))
		blt.Refresh()
		key := blt.Read()

		// Clear each entity off the screen, so we can re-draw them
		for _, e := range entities {
			e.Clear()
		}

		if key != blt.TK_CLOSE {
			handleInput(key)
		} else {
			break
		}

		renderAll()

	}

	blt.Close()
}

func saveState() {
	runtime.LockOSThread()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(gameMap)
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

func loadState() {
	runtime.LockOSThread()
	fmt.Println("LOAD START")
	b, err := ioutil.ReadFile("save.txt")
	if err != nil {
		fmt.Println("PANIC LOAD")
		panic(err)
	}

	buf := bytes.NewBuffer(b)

	dec := gob.NewDecoder(buf)

	var gameMap2 ui.GameMap
	err = dec.Decode(&gameMap2)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Println("LOAD END")
	runtime.UnlockOSThread()
	gameMap = &gameMap2
}

func handleInput(key int) {
	var (
		dx, dy int
	)

	switch key {
	case blt.TK_RIGHT:
		dx, dy = 1, 0
	case blt.TK_LEFT:
		dx, dy = -1, 0
	case blt.TK_UP:
		dx, dy = 0, -1
	case blt.TK_DOWN:
		dx, dy = 0, 1
	case blt.TK_S:
		saveState()
	case blt.TK_L:
		loadState()
	}

	player.Move(dx, dy)
}

func renderEntities() {
	for _, e := range entities {
		e.Draw()
	}
}

func renderMap() {
	for x := 0; x < gameMap.Width; x++ {
		for y := 0; y < gameMap.Height; y++ {
			if gameMap.Tiles[x][y].Blocked == true {
				blt.Color(blt.ColorFromName("gray"))
				blt.Print(x, y, "#")
			} else {
				blt.Color(blt.ColorFromName("brown"))
				blt.Print(x, y, ".")
			}
		}
	}
}

func renderAll() {
	renderMap()
	renderEntities()
}
