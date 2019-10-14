package ui

import (
	blt "bearlibterminal"
	"strconv"
)

type GameEntity struct {
	X     int
	Y     int
	Layer int
	Char  string
	Color string
}

func (e *GameEntity) Move(dx int, dy int) {
	e.X += dx
	e.Y += dy
}

func (e *GameEntity) Draw() {
	blt.Layer(e.Layer)
	blt.Color(blt.ColorFromName(e.Color))
	blt.Print(e.X, e.Y, e.Char)
}

func (e *GameEntity) Clear() {
	blt.Layer(e.Layer)
	blt.Print(e.X, e.Y, " ")
}

type Tile struct {
	Blocked     bool
	BlocksSight bool
	Baba1       string
	Baba2       string
	Baba3       string
	Baba4       string
	Baba5       string
	Baba6       string
	Baba7       string
	Baba8       string
	Baba9       string
}

type GameMap struct {
	Width  int
	Height int
	Tiles  [][]*Tile
}

func (m *GameMap) InitializeMap() {
	m.Tiles = make([][]*Tile, m.Width)
	for i := range m.Tiles {
		m.Tiles[i] = make([]*Tile, m.Height)
	}

	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			if x == 0 || x == m.Width-1 || y == 0 || y == m.Height-1 {
				m.Tiles[x][y] = &Tile{true, true, "hahahahahahahahahhahahahah" + strconv.Itoa(x+y), "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah"}
			} else {
				m.Tiles[x][y] = &Tile{false, false, "hahahahahahahahahhahahahah" + strconv.Itoa(x+y), "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah", "hahahahahahahahahhahahahah"}
			}
		}
	}
}

func (m *GameMap) IsBlocked(x int, y int) bool {
	if m.Tiles[x][y].Blocked {
		return true
	} else {
		return false
	}
}
