package ui

import (
	blt "bearlibterminal"

	m "github.com/castle/src/game/model"
)

func renderMap(camera *Camera, world *m.World) {
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
			if tiles[rowIndex][columnIndex].Surface == m.SurfaceRock {
				blt.Color(blt.ColorFromName("gray"))
				blt.Print(x*TileSizeX, y*TileSizeY, "[font=tile]#[/font]")
			} else {
				blt.Color(blt.ColorFromName("brown"))
				blt.Print(x*TileSizeX, y*TileSizeY, "[font=tile].[/font]")
			}
		}
	}
}

func renderPlayer(camera *Camera, player *m.Player) {
	camX := int(camera.Width / TileSizeX)
	camY := int(camera.Height / TileSizeY)
	x := camX + player.Pos.X - camera.Pos.X
	y := camY + player.Pos.Y - camera.Pos.Y
	blt.Color(blt.ColorFromName("white"))
	blt.Print(x*TileSizeX, y*TileSizeY, "[font=tile]@[/font]")
}
