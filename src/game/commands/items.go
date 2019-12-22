package commands

import (
	m "github.com/castle/src/game/model"
	"github.com/castle/src/game/world"
)

func GetFoodAroundPlayer(gs *m.State) []m.Food {
	var pos = &gs.Characters[0].Pos
	return GetFoodAround(gs, pos)
}

func GetFoodAround(gs *m.State, pos *m.Pos) []m.Food {
	var food []m.Food
	var tiles = world.GetTileAroundPos(gs, pos)
	for _, tile := range tiles {
		food = append(food, tile.Items.Food...)
	}
	return food
}

func GetLiquidTilesAroundPlayer(gs *m.State) []m.TileAndPos {
	var pos = &gs.Characters[0].Pos
	var liquidTiles []m.TileAndPos
	var tiles = world.GetTileAndPosAroundPos(gs, pos)
	for _, tile := range tiles {
		if tile.Tile.Surface == m.SurfaceWater || tile.Tile.Volume == m.VolumeWater {
			liquidTiles = append(liquidTiles, tile)
		}
	}
	return liquidTiles
}
