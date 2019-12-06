package character

import m "github.com/castle/src/game/model"

func GetTile(gs *m.State, characterID int) *m.Tile {
	pos := gs.Characters[characterID].Pos
	return &gs.World.Regions[pos.Region].Tiles[pos.Z][pos.X][pos.Y]
}
