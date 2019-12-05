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
			if tile.Surface == m.SurfaceRock {
				blt.Color(blt.ColorFromName("gray"))
				blt.Print(x*TileSizeX, y*TileSizeY, "[font=tile]#[/font]")
			} else {
				blt.Color(blt.ColorFromName("brown"))
				blt.Print(x*TileSizeX, y*TileSizeY, "[font=tile].[/font]")
			}
			if len(tile.Items.Food) > 0 {
				// blt.Composition(blt.TK_ON)
				// blt.Layer(1)
				blt.Color(blt.ColorFromName(Colors[ColorRed]))
				blt.Print(x*TileSizeX, y*TileSizeY, "p")
				// blt.Layer(0)
				// blt.Composition(blt.TK_OFF)
			}
		}
	}
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
