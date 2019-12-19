package item

import (
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func UpdateStackQuantity(gs *m.State, quantity int, tile *m.Tile, characterID int, where int, index int, itemType int) {
	switch itemType {
	case c.ItemTypeFood:
		var slice []m.Food
		var sliceP *[]m.Food
		switch where {
		case c.WhereFloor:
			slice = tile.Items.Food
			sliceP = &tile.Items.Food
		case c.WhereInventory:
			slice = gs.Characters[characterID].Inventory.Food
			sliceP = &gs.Characters[characterID].Inventory.Food
		}
		slice[index].Quantity += quantity
		if slice[index].Quantity <= 0 {
			if len(slice) == 1 {
				*sliceP = nil
			} else {
				slice[index] = slice[len(slice)-1]
				*sliceP = slice[:len(slice)-1]
			}
		}
	}
}

func AddFoodToStack(gs *m.State, food m.Food, tile *m.Tile, characterID int, where int) {
	switch where {
	case c.WhereInventory:
		inventory := &gs.Characters[characterID].Inventory
		inventory.Food = append(inventory.Food, food)
	}
}
