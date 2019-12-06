package ui

import (
	blt "bearlibterminal"
	"fmt"
	"strconv"

	cmd "github.com/castle/src/game/commands"
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

const (
	InfoPanelLeftMargin = 1
	InfoPanelTopMargin  = 1
)

func renderInfoPanel(texts []UiElement, buttons []UiElement) {

	blt.BkColor(blt.ColorFromName(Colors[ColorWhitish]))
	blt.ClearArea(InforPanelStart, 0, InfoPanelDefaultWidth, CameraDefaultHeight*TileSizeY)

	renderElements(texts)
	renderElements(buttons)
}

func setInfoPanel(ui *State, gs *m.State) {
	nextRow := InfoPanelTopMargin

	if ui.IntendedAction > 0 {
		// possible actions
		setActionsDetails(ui, gs, &nextRow)
	} else {
		// region name
		region := &gs.World.Regions[ui.Camera.Pos.Region]
		text := region.Name
		action := Action{Name: "toggleInfoDetails", EntityType: EntityTypeRegion, Entity: ui.Camera.Pos.Region}
		addElementToInfoPanel(ui, text, &nextRow, InfoPanelLeftMargin, 0, action)
		// time
		text = getTimeString(gs.Time)
		addElementToInfoPanel(ui, text, &nextRow, InfoPanelLeftMargin, 0, action)
		// possible actions
		setActionsButtons(ui, gs, &nextRow)
		// player details
		setPlayerDetails(ui, gs.Characters[0], &nextRow)
		// region details
		if ui.EntityDetails.Type == EntityTypeRegion {
			setInfoPanelRegionDetails(ui, region, &nextRow)
		}
		// tile details
		tileDetailsX := ui.Camera.Pos.X
		tileDetailsY := ui.Camera.Pos.Y
		if ui.EntityDetails.Type == EntityTypeTile {
			tileDetailsX = ui.EntityDetails.Data1
			tileDetailsY = ui.EntityDetails.Data2
		}
		setInfoPanelTileDetails(ui, region, tileDetailsX, tileDetailsY, &nextRow)
	}

	// command log
	setInfoPanelLog(ui, gs, &nextRow)
}

func addElementToInfoPanel(ui *State, text string, nextRow *int, offset int, color int, leftClick Action) {
	w, h := blt.MeasureExt(InfoPanelDefaultWidth*TextSizeX, 100, text)
	width := w * TextSizeX
	height := h * TextSizeY
	element := UiElement{
		X:           InforPanelStart + offset,
		Y:           *nextRow,
		Height:      height,
		Width:       width,
		Color:       color,
		Text:        text,
		OnLeftClick: leftClick,
	}
	if leftClick.Name != "" {
		ui.Buttons = append(ui.Buttons, element)
	} else {
		ui.Texts = append(ui.Texts, element)
	}
	*nextRow += height
}

func addVerticalMargin(nextRow *int) {
	*nextRow++
}

func getFoodLabel(food m.Food) string {
	return fmt.Sprintf("%d %ss", food.Quantity, m.FoodSubtypeNames[food.Subtype])
}

func setInfoPanelRegionDetails(ui *State, region *m.Region, nextRow *int) {
	text := region.Description
	addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, Action{})
}

func setInfoPanelTileDetails(ui *State, region *m.Region, x int, y int, nextRow *int) {
	addVerticalMargin(nextRow)
	tile := region.Tiles[ui.Camera.Pos.Z][x][y]
	text := m.SurfaceNames[tile.Surface]
	addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, Action{})
	if len(tile.Items.Food) > 0 {
		for _, food := range tile.Items.Food {
			text := getFoodLabel(food)
			addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, Action{})
		}
	}
}

func setInfoPanelLog(ui *State, gs *m.State, nextRow *int) {
	for _, logVal := range gs.Log.Logs {
		text := logVal.Text
		color := ColorBlackish
		if logVal.LogType == m.LogTypeAlert {
			color = ColorRed
			text = "ACHTUNG! " + text
		}
		addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, color, Action{})
	}
}

func setPlayerDetails(ui *State, player *m.Character, nextRow *int) {
	text := "Faim: " + strconv.Itoa(player.Needs.Hunger)
	addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, Action{})
	text = "Nutrition - Total: " + strconv.Itoa(player.Physical.Nutrition.Total)
	addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, Action{})
	if player.Physical.Alive == false {
		text = "MOURRU!"
		addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, ColorRed, Action{})
	}
}

func setActionsButtons(ui *State, gs *m.State, nextRow *int) {
	for _, action := range BasicPlayerActions {
		text := "(" + action.Handle + ") " + action.Name
		addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, Action{})
	}
}

func setActionsDetails(ui *State, gs *m.State, nextRow *int) {
	switch ui.IntendedAction {
	case ActionEat:
		setEat(ui, gs, nextRow)
	case ActionUseInventory:
		setUseInventory(ui, gs, nextRow)
	case ActionPickup:
		setPickup(ui, gs, nextRow)
	}
}

func setEat(ui *State, gs *m.State, nextRow *int) {
	addElementToInfoPanel(ui, "MANGER", nextRow, InfoPanelLeftMargin, 0, Action{})
	foodOnGround := cmd.GetPlayerTile(gs).Items.Food
	if len(foodOnGround) > 0 {
		addElementToInfoPanel(ui, "Au Sol", nextRow, InfoPanelLeftMargin, 0, Action{})
		for i, food := range foodOnGround {
			addElementToInfoPanel(ui, getFoodLabel(food), nextRow, InfoPanelLeftMargin, 0, Action{
				Name:   "eat",
				Entity: i,
				Data:   c.WhereFloor,
			})
		}
	}
	foodInInventory := gs.Characters[0].Inventory.Food
	if len(foodInInventory) > 0 {
		addElementToInfoPanel(ui, "Dans l'inventaire", nextRow, InfoPanelLeftMargin, 0, Action{})
		for i, food := range foodInInventory {
			addElementToInfoPanel(ui, getFoodLabel(food), nextRow, InfoPanelLeftMargin, 0, Action{
				Name:   "eat",
				Entity: i,
				Data:   c.WhereInventory,
			})
		}
	}
}

func setUseInventory(ui *State, gs *m.State, nextRow *int) {
	addElementToInfoPanel(ui, "INVENTAIRE", nextRow, InfoPanelLeftMargin, 0, Action{})
	inventory := gs.Characters[0].Inventory
	for _, food := range inventory.Food {
		text := getFoodLabel(food)
		addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, Action{})
	}
}

func setPickup(ui *State, gs *m.State, nextRow *int) {
	addElementToInfoPanel(ui, "RAMASSER", nextRow, InfoPanelLeftMargin, 0, Action{})
	foodOnGround := cmd.GetPlayerTile(gs).Items.Food
	if len(foodOnGround) > 0 {
		for i, food := range foodOnGround {
			text := getFoodLabel(food)
			addElementToInfoPanel(ui, text, nextRow, InfoPanelLeftMargin, 0, Action{
				Name:       "pickup",
				Entity:     i,
				EntityType: c.ItemTypeFood,
				Data:       c.WhereFloor,
			})
		}
	}
}
