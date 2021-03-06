package world

import (
	"math/rand"

	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func GetTileAtPos(gs *m.State, pos m.Pos) *m.Tile {
	return &gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y]
}

func GetTileAroundPos(gs *m.State, pos *m.Pos) []m.Tile {
	return []m.Tile{
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y+1],
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y-1],
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y],
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X+1][pos.Y+1],
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X+1][pos.Y-1],
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X+1][pos.Y],
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X-1][pos.Y+1],
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X-1][pos.Y-1],
		gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X-1][pos.Y],
	}
}

func GetTileAndPosAroundPos(gs *m.State, pos *m.Pos) []m.TileAndPos {
	return []m.TileAndPos{
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y+1], Pos: m.Pos{pos.Region, pos.X, pos.Y + 1, pos.Z}},
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y-1], Pos: m.Pos{pos.Region, pos.X, pos.Y - 1, pos.Z}},
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y], Pos: m.Pos{pos.Region, pos.X, pos.Y, pos.Z}},
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X+1][pos.Y+1], Pos: m.Pos{pos.Region, pos.X + 1, pos.Y + 1, pos.Z}},
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X+1][pos.Y-1], Pos: m.Pos{pos.Region, pos.X + 1, pos.Y - 1, pos.Z}},
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X+1][pos.Y], Pos: m.Pos{pos.Region, pos.X + 1, pos.Y, pos.Z}},
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X-1][pos.Y+1], Pos: m.Pos{pos.Region, pos.X - 1, pos.Y + 1, pos.Z}},
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X-1][pos.Y-1], Pos: m.Pos{pos.Region, pos.X - 1, pos.Y - 1, pos.Z}},
		{Tile: gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X-1][pos.Y], Pos: m.Pos{pos.Region, pos.X - 1, pos.Y, pos.Z}},
	}
}

func GetInitialWorld() *m.World {

	regions := make([]m.Region, c.WorldSize)

	for i := 0; i < c.WorldSize; i++ {

		tiles := make([][]m.Tile, c.RegionWidth)
		for i := range tiles {
			tiles[i] = make([]m.Tile, c.RegionHeight)
		}

		for x := 0; x < c.RegionWidth; x++ {
			for y := 0; y < c.RegionHeight; y++ {
				if x == 0 || x == c.RegionWidth-1 || y == 0 || y == c.RegionHeight-1 {
					tiles[x][y] = m.Tile{Surface: m.SurfaceRock, Volume: m.VolumeRock}
				} else {
					tiles[x][y] = m.Tile{Surface: m.SurfaceGround, Volume: m.VolumeAir}
					r := rand.Intn(100)
					if r > 80 {
						tiles[x][y].Items = m.TileItems{
							Food: []m.Food{{Type: 0, Subtype: 0, Quantity: 1, Nutrition: 1}, {Type: 0, Subtype: 1, Quantity: 2, Nutrition: 1}},
						}
					}
				}
			}
		}

		plane := make([][][]m.Tile, 1)
		plane[0] = tiles

		regions[i] = m.Region{
			Name: "A beautiful region",
			Description: `
			The Vosges department is one of the original 83 departments of France, created on February 9, 1790 during the French Revolution.[4] It was made of territories that had been part of the province of Lorraine. In German it is referred to as Vogesen.

			In 1793 the independent principality of Salm-Salm (town of Senones and its surroundings), enclosed inside the Vosges department, was annexed to France and incorporated into Vosges. In 1795 the area of Schirmeck was detached from the Bas-Rhin department and incorporated into the Vosges department.[5] The Vosges department had now an area of 6,127 km² (2,366 sq. miles) which it kept until 1871.

			In 1794 the Vosges was the site of a major battle between the forces of Revolutionary France and the Allied Coalition. See Battle of Trippstadt.

			The Place des Vosges in Paris was so renamed in 1799 when the department became the first to pay the new Revolutionary taxes.
			`,
			Tiles: plane,
		}

		if i == 0 {
			MakeRiver(&regions[i])
		}
	}

	return &m.World{
		Regions: regions,
	}
}
