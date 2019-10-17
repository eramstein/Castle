package game

type World struct {
	Regions []*Region
}

type Region struct {
	Name        string
	Description string
	Tiles       [][][]*Tile
}

type Tile struct {
	Surface int
}

const (
	SurfaceGround = 1
	SurfaceRock   = 2
)

var SurfaceNames = map[int]string{
	1: "Ground",
	2: "Rock",
}

func InitWorld() *World {

	regions := make([]*Region, WorldSize)

	for i := 0; i < WorldSize; i++ {

		tiles := make([][]*Tile, RegionWidth)
		for i := range tiles {
			tiles[i] = make([]*Tile, RegionHeight)
		}

		for x := 0; x < RegionWidth; x++ {
			for y := 0; y < RegionHeight; y++ {
				if x == 0 || x == RegionWidth-1 || y == 0 || y == RegionHeight-1 {
					tiles[x][y] = &Tile{Surface: SurfaceRock}
				} else {
					tiles[x][y] = &Tile{Surface: SurfaceGround}
				}
			}
		}

		plane := make([][][]*Tile, 1)
		plane[0] = tiles

		regions[i] = &Region{
			Tiles: plane,
		}
	}

	return &World{
		Regions: regions,
	}
}
