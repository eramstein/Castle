package ui

const (
	InfoPanelDefaultWidth = 30
	CameraDefaultWidth    = 61
	CameraDefaultHeight   = 61
	InforPanelStart       = CameraDefaultWidth * 2
)

const (
	ColorWhite = iota
	ColorRed   = iota
)

var Colors = map[int]string{
	ColorWhite: "#FFFFFF",
	ColorRed:   "#FF0000",
}
