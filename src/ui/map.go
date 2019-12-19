package ui

import (
	blt "bearlibterminal"

	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func renderMap(camera Camera, world *m.World) {
	tiles := world.Regions[camera.Pos.Region].Tiles[camera.Pos.Z]
	for x := 0; x < camera.Width; x++ {
		columnIndex := camera.Pos.X - int(camera.Width/TileSizeX) + x
		if columnIndex < 0 || columnIndex >= len(tiles) {
			continue
		}
		for y := 0; y < camera.Height; y++ {
			column := tiles[columnIndex]
			rowIndex := camera.Pos.Y - int(camera.Height/TileSizeY) + y
			if rowIndex < 0 || rowIndex >= len(column) {
				continue
			}
			tile := &tiles[columnIndex][rowIndex]
			renderSurface(tile.Surface, x, y)
			if len(tile.Items.Food) > 0 {
				// blt.Composition(blt.TK_ON)
				// blt.Layer(1)
				blt.Color(blt.ColorFromName(Colors[ColorRed]))
				blt.Print(x*TileSizeX, y*TileSizeY, "[font=tile]p[/font]")
				// blt.Layer(0)
				// blt.Composition(blt.TK_OFF)
			}
		}
	}
}

func renderSurface(surface, x, y int) {
	switch surface {
	case m.SurfaceRock:
		blt.BkColor(blt.ColorFromName("gray"))
	default:
		blt.BkColor(blt.ColorFromName("black"))
	}
	blt.ClearArea(x*TileSizeX, y*TileSizeY, 2, 2)
}

func renderPlayer(camera Camera, player *m.Character) {
	camX := int(camera.Width / TileSizeX)
	camY := int(camera.Height / TileSizeY)
	x := camX + player.Pos.X - camera.Pos.X
	y := camY + player.Pos.Y - camera.Pos.Y
	blt.Color(blt.ColorFromName("white"))
	blt.Print(x*TileSizeX, y*TileSizeY, "[font=tile]@[/font]")
}

func tileAtCoordinates(camera Camera, x, y int) (bool, int, int) {
	tileOnScreenX := int(x / TileSizeX)
	if tileOnScreenX >= CameraDefaultWidth {
		return false, 0, 0
	}
	tileDiffCameraX := tileOnScreenX - int(camera.Width/2)
	tileOnScreenY := int(y / TileSizeY)
	tileDiffCameraY := tileOnScreenY - int(camera.Height/2)

	tileX := camera.Pos.X + tileDiffCameraX
	tileY := camera.Pos.Y + tileDiffCameraY

	if tileX < 0 || tileX >= c.RegionWidth || tileY < 0 || tileY >= c.RegionHeight {
		return false, tileX, tileY
	}

	return true, tileX, tileY
}
