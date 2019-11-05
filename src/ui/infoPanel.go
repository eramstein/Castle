package ui

import (
	blt "bearlibterminal"

	m "github.com/castle/src/game/model"
)

const (
	InfoPanelLeftMargin = 1
	InfoPanelTopMargin  = 1
)

func renderInfoPanel(texts []*UiElement, buttons []*UiElement) {

	blt.BkColor(blt.ColorFromName(Colors[ColorWhitish]))
	blt.ClearArea(InforPanelStart, 0, InfoPanelDefaultWidth, CameraDefaultHeight*TileSizeY)

	renderElements(texts)
	renderElements(buttons)
}

func setInfoPanel(ui *State, gs *m.State) {
	nextRow := InfoPanelTopMargin

	region := gs.World.Regions[ui.Camera.Pos.Region]

	// region name
	text := region.Name
	action := &Action{Name: "toggleInfoDetails", EntityType: EntityTypeRegion, Entity: ui.Camera.Pos.Region}
	addElementToInfoPanel(ui, text, &nextRow, InfoPanelLeftMargin, 0, action)
	// time

	// region details
	if ui.EntityDetails.Type == EntityTypeRegion {
		setInfoPanelRegionDetails(ui, region, &nextRow)
	}
}

func addElementToInfoPanel(ui *State, text string, nextRow *int, offset int, color int, leftClick *Action) {
	w, h := blt.MeasureExt(InfoPanelDefaultWidth*TextSizeX, 100, text)
	width := w * TextSizeX
	height := h * TextSizeY
	element := &UiElement{
		X:           InforPanelStart + offset,
		Y:           *nextRow,
		Height:      height,
		Width:       width,
		Color:       color,
		Text:        text,
		OnLeftClick: leftClick,
	}
	if leftClick != nil {
		ui.Buttons = append(ui.Buttons, element)
	} else {
		ui.Texts = append(ui.Texts, element)
	}
	*nextRow += height
}

func setInfoPanelRegionDetails(ui *State, region *m.Region, nextRow *int) {
	text := region.Description
	addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, nil)
}
