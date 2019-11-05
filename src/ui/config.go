package ui

import "strconv"

const (
	Title                 = "Château Ramstein"
	Font                  = "UbuntuMono.ttf"
	CellPixelSize         = 10
	TileSizeX             = 2
	TileSizeY             = 2
	TextSizeX             = 1
	TextSizeY             = 3
	InfoPanelDefaultWidth = 60
	CameraDefaultWidth    = 61
	CameraDefaultHeight   = 61
	InforPanelStart       = CameraDefaultWidth * TileSizeX
	WindowSizeX           = CameraDefaultWidth*TileSizeX + InfoPanelDefaultWidth*TextSizeX
	WindowSizeY           = CameraDefaultHeight * TileSizeY
)

const (
	ColorWhite    = iota
	ColorRed      = iota
	ColorWhitish  = iota
	ColorBlackish = iota
)

var Colors = map[int]string{
	ColorWhite:    "#FFFFFF",
	ColorWhitish:  "#FCFCFC",
	ColorBlackish: "#333333",
	ColorRed:      "#FF0000",
}

var size = "size=" + strconv.Itoa(WindowSizeX) + "x" + strconv.Itoa(WindowSizeY)
var title = "title='" + Title + "'"
var cellsize = "cellsize=" + strconv.Itoa(CellPixelSize) + "x" + strconv.Itoa(CellPixelSize)
var window = "window: " + size + "," + title + "," + cellsize
var input = "input.filter = [keyboard, mouse]"
var textFont = "text font: " + Font + ",size=16,spacing=1x3"
var tileFont = "tile font: " + Font + ",size=18"

var BltConfig = window + "; " + input + "; " + textFont + "; " + tileFont + ";"
