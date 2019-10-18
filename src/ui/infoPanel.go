package ui

import (
	blt "bearlibterminal"
)

func renderInfoPanel(texts []*UiElement, buttons []*UiElement) {
	xStart := InforPanelStart
	xEnd := xStart + InfoPanelDefaultWidth*2
	yEnd := CameraDefaultHeight * 3

	blt.BkColor(blt.ColorFromName("blue"))
	for x := xStart; x < xEnd; x++ {
		for y := 0; y < yEnd; y++ {
			blt.Print(x, y, " ")
		}
	}
	datext := "[font=text]Yep, this is exactly what I do. The base tile size is the smallest of all the cells that I use and everything else fits with some multiple of that base cell. It lets me use small tiles for status bars, large tiles for maps, different sized ascii for maps as well as taller text to help with readability.[/font]"
	_, h := blt.MeasureExt(InfoPanelDefaultWidth*2, 100, datext)
	blt.PrintExt(xStart+1, 1, InfoPanelDefaultWidth*2, h, blt.TK_ALIGN_MIDDLE, datext)
	blt.BkColor(0)

	renderElements(texts)
	renderElements(buttons)
}
