package character

import (
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
	"github.com/davecgh/go-spew/spew"
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
		spew.Dump(gs.Player.Pos)
	} else {
		gs.Characters[characterID].Pos = pos
		// TODO: create new move task from path finding
	}
}
