package model

type World struct {
	Regions []Region
}

type Region struct {
	Name        string
	Description string
	Tiles       [][][]Tile
	Rivers      []River
}

type TileAndPos struct {
	Tile Tile
	Pos  Pos
}

type Tile struct {
	Surface      int
	Volume       int
	SurfaceDepth int
	Items        TileItems
	River        *River
	Zones        []Zone
}

type Zone struct {
	Type int
}

type TileItems struct {
	Food []Food
}

type Pos struct {
	Region int
	X      int
	Y      int
	Z      int
}

type River struct {
	Name             string
	InDirection      int
	OutDirection     int
	CurrentStrength  int
	WaterTemperature int
	Depth            int
}

const (
	DirectionNorth = iota
	DirectionSouth = iota
	DirectionEast  = iota
	DirectionWest  = iota
)

const (
	SurfaceGround = 0
	SurfaceRock   = 1
	SurfaceWater  = 2
)

var SurfaceNames = map[int]string{
	0: "Ground",
	1: "Rock",
	2: "Water",
}

const (
	VolumeAir   = 0
	VolumeRock  = 1
	VolumeWater = 2
)

var VolumeNames = map[int]string{
	0: "Air",
	1: "Rock",
	2: "Water",
}

var VolumeBlocking = map[int]bool{
	0: false,
	1: true,
}
