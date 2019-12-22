package world

import (
	"fmt"
	"math/rand"

	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
	"github.com/castle/src/game/utils"
)

type inflexion struct {
	x         int
	y         int
	inLength  int
	outLength int
}

type coord struct {
	x int
	y int
}

func MakeRiver(region *m.Region) {
	defer utils.Track(utils.Runningtime("River creation"))

	var river m.River
	var fromDirection int
	var toDirection int
	var inPos int
	var outPos int
	var fromBorderLength int
	var outBorderLength int
	var riverWidth int
	var inflexionPoint inflexion
	var inflexionsCount int
	var startCoord coord

	riverWidth = 3 + rand.Intn(4)
	fromDirection = rand.Intn(3)
	toDirection = rand.Intn(3)

	river.Name = "Le Rhin"
	river.InDirection = fromDirection
	river.OutDirection = toDirection
	river.WaterTemperature = 20
	river.CurrentStrength = 5
	river.Depth = 8

	if fromDirection == toDirection {
		fromDirection = (toDirection + 1) % 3
	}
	if fromDirection == m.DirectionNorth || fromDirection == m.DirectionSouth {
		fromBorderLength = c.RegionWidth
	} else {
		fromBorderLength = c.RegionHeight
	}
	inPos = rand.Intn(fromBorderLength - riverWidth)
	if toDirection == m.DirectionNorth || toDirection == m.DirectionSouth {
		outBorderLength = c.RegionWidth
	} else {
		outBorderLength = c.RegionHeight
	}
	outPos = rand.Intn(outBorderLength - riverWidth)

	if fromDirection == m.DirectionNorth && toDirection != m.DirectionSouth ||
		fromDirection == m.DirectionSouth && toDirection != m.DirectionNorth ||
		fromDirection == m.DirectionEast && toDirection != m.DirectionWest ||
		fromDirection == m.DirectionWest && toDirection != m.DirectionEast {
		inflexionsCount = 1
		inflexionPoint = inflexion{80 + rand.Intn(c.RegionWidth-80), 80 + rand.Intn(c.RegionHeight-80), 0, 0}
	} else {
		inflexionsCount = 0
	}

	switch fromDirection {
	case m.DirectionNorth:
		startCoord = coord{inPos, 0}
	case m.DirectionSouth:
		startCoord = coord{inPos, c.RegionHeight}
	case m.DirectionEast:
		startCoord = coord{c.RegionWidth, inPos}
	case m.DirectionWest:
		startCoord = coord{0, inPos}
	}

	if inflexionsCount == 0 {
		buildRiverSection(region, startCoord, fromDirection, outPos, inPos, fromBorderLength, riverWidth, &river)
	}
	if inflexionsCount == 1 {
		// treat inflexion point as the out position
		var intermediateOutPos int
		var inflexionPointDistance int
		var oppositeDirection int
		switch fromDirection {
		case m.DirectionNorth:
			intermediateOutPos = inflexionPoint.x
			inflexionPointDistance = inflexionPoint.y
		case m.DirectionSouth:
			intermediateOutPos = inflexionPoint.x
			inflexionPointDistance = c.RegionHeight - inflexionPoint.y
		case m.DirectionEast:
			intermediateOutPos = inflexionPoint.y
			inflexionPointDistance = c.RegionWidth - inflexionPoint.x
		case m.DirectionWest:
			intermediateOutPos = inflexionPoint.y
			inflexionPointDistance = inflexionPoint.x
		}
		buildRiverSection(region, startCoord, fromDirection, intermediateOutPos, inPos, inflexionPointDistance, riverWidth, &river)
		// turn around and go to the actual out position
		var shiftBeforeturnX = 0
		var shiftBeforeturnY = 0
		var remainingDistance int
		var intermediateInPos int
		switch toDirection {
		case m.DirectionNorth:
			oppositeDirection = m.DirectionSouth
			remainingDistance = inflexionPoint.y
			intermediateInPos = inflexionPoint.x
		case m.DirectionSouth:
			oppositeDirection = m.DirectionNorth
			remainingDistance = c.RegionHeight - inflexionPoint.y
			intermediateInPos = inflexionPoint.x
		case m.DirectionEast:
			oppositeDirection = m.DirectionWest
			remainingDistance = c.RegionWidth - inflexionPoint.x
			intermediateInPos = inflexionPoint.y
		case m.DirectionWest:
			oppositeDirection = m.DirectionEast
			remainingDistance = inflexionPoint.x
			intermediateInPos = inflexionPoint.y
		}
		// shift a bit the turning point in some cases to smoothen the angle
		switch fromDirection {
		case m.DirectionNorth:
			if outPos < inflexionPoint.x {
				shiftBeforeturnY = -riverWidth / 2
			}
		case m.DirectionSouth:
			if outPos < inflexionPoint.x {
				shiftBeforeturnY = riverWidth / 2
			}
		case m.DirectionEast:
			if outPos < inflexionPoint.y {
				shiftBeforeturnX = riverWidth / 2
			}
		case m.DirectionWest:
			if outPos < inflexionPoint.y {
				shiftBeforeturnX = -riverWidth / 2
			}
		}
		// TODO: handle real turn instead of shifting
		fmt.Println("shiftBeforeturnY", shiftBeforeturnY)
		buildRiverSection(region, coord{inflexionPoint.x + shiftBeforeturnX, inflexionPoint.y + shiftBeforeturnY}, oppositeDirection, outPos, intermediateInPos, remainingDistance, riverWidth, &river)
	}

}

func buildRiverSection(region *m.Region, startCoord coord, fromDirection, targetPos, inPos, sectionLength, riverWidth int, river *m.River) {
	var shiftsDone = 0
	var padding = 0

	var shiftsToDo = targetPos - inPos
	shiftOpposite := false
	if targetPos < inPos {
		shiftOpposite = true
		shiftsToDo = -shiftsToDo
	}

	// TODO: sinuosités en changeant target pos en cours de route (mais même direction})

	fmt.Println("startCoord", startCoord)
	fmt.Println("fromDirection", fromDirection)
	fmt.Println("targetPos", targetPos)
	fmt.Println("inPos", inPos)
	fmt.Println("sectionLength", sectionLength)
	fmt.Println("riverWidth", riverWidth)

	for l := 0; l < sectionLength; l++ {
		padding = 0
		var shiftProba float32
		if shiftsToDo-shiftsDone >= sectionLength-l {
			shiftsDone += (shiftsToDo-shiftsDone)/(sectionLength-l) + 1
		} else {
			shiftProba = float32(shiftsToDo-shiftsDone) / float32(sectionLength-l)
			if rand.Float32() <= shiftProba {
				shiftsDone++
				padding = 1
			}
		}

		shiftVal := shiftsDone
		start := -padding
		end := riverWidth

		if shiftsToDo-shiftsDone >= sectionLength-l {
			end += riverWidth * (shiftsToDo - shiftsDone) / (sectionLength - l)
		}

		if shiftOpposite {
			shiftVal = -shiftVal
			start = 0
			end = riverWidth + 1
		}

		for w := start; w < end; w++ {
			switch fromDirection {
			case m.DirectionNorth:
				fillTile(region, river, coord{startCoord.x + w + shiftVal, startCoord.y + l})
			case m.DirectionSouth:
				fillTile(region, river, coord{startCoord.x + w + shiftVal, startCoord.y - l})
			case m.DirectionWest:
				fillTile(region, river, coord{startCoord.x + l, startCoord.y + w + shiftVal})
			case m.DirectionEast:
				fillTile(region, river, coord{startCoord.x - l, startCoord.y + w + shiftVal})
			}
		}
	}
}

func fillTile(region *m.Region, river *m.River, pos coord) {
	if pos.x >= 0 && pos.y >= 0 && pos.x < c.RegionWidth && pos.y < c.RegionHeight {
		tile := &region.Tiles[0][pos.x][pos.y]
		tile.Surface = m.SurfaceWater
		tile.SurfaceDepth = river.Depth
		tile.River = river
	}
}
