package character

import (
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func GetTimeMoveToTile(atts m.Attributes, physicalState m.PhysicalState, tile m.Tile) (seconds int) {
	return 10
}

func MoveCharacter(gs *m.State, characterID int, pos m.Pos) {
	if characterID < c.PlayerAgentID {
		return
	}

	if characterID == c.PlayerAgentID {
		gs.Player.Pos = pos
	} else {
		gs.Characters[characterID].Pos = pos
		// TODO: create new move task from path finding
	}
}

func IsMoveValid(gs *m.State, pos m.Pos) bool {
	if pos.X < 0 || pos.X >= c.RegionWidth || pos.Y < 0 || pos.Y >= c.RegionHeight {
		return false
	}
	tile := gs.World.Regions[pos.Region].Tiles[pos.Z][pos.Y][pos.X]
	if m.VolumeBlocking[tile.Volume] {
		return false
	}
	return true
}
