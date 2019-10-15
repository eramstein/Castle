package sim

import (
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func InitWorld() *m.World {

	regions := make([]*m.Region, c.WorldSize)

	for i := 0; i < c.WorldSize; i++ {

		tiles := make([][]*m.Tile, c.RegionWidth)
		for i := range tiles {
			tiles[i] = make([]*m.Tile, c.RegionHeight)
		}

		for x := 0; x < c.RegionWidth; x++ {
			for y := 0; y < c.RegionHeight; y++ {
				if x == 0 || x == c.RegionWidth-1 || y == 0 || y == c.RegionHeight-1 {
					tiles[x][y] = &m.Tile{Surface: m.SurfaceRock}
				} else {
					tiles[x][y] = &m.Tile{Surface: m.SurfaceGround}
				}
			}
		}

		plane := make([][][]*m.Tile, 1)
		plane[0] = tiles

		regions[i] = &m.Region{
			Tiles: plane,
		}
	}

	return &m.World{
		Regions: regions,
	}
}
