package model

type World struct {
	Regions []Region
}

type Region struct {
	Name        string
	Description string
	Tiles       [][][]Tile
}

type Tile struct {
	Surface int
	Volume  int
	Items   TileItems
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
	inDirection      int
	outDirection     int
	currentStrength  int
	waterTemperature int
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
	VolumeAir  = 0
	VolumeRock = 1
)

var VolumeNames = map[int]string{
	0: "Air",
	1: "Rock",
}

var VolumeBlocking = map[int]bool{
	0: false,
	1: true,
}
