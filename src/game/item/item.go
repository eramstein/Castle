package item

import (
	c "github.com/castle/src/game/config"
	m "github.com/castle/src/game/model"
)

func UpdateStackQuantity(gs *m.State, quantity int, tile *m.Tile, characterID int, where int, index int, itemType int) {
	switch where {
	case c.WhereFloor:
		slice := tile.Items.Food
		slice[index].Quantity += quantity
		if slice[index].Quantity <= 0 {
			if len(slice) == 1 {
				tile.Items.Food = nil
			} else {
				slice[index] = slice[len(slice)-1]
				tile.Items.Food = slice[:len(slice)-1]
			}
		}
	case c.WhereInventory:
		slice := gs.Characters[characterID].Inventory.Food
		slice[index].Quantity += quantity
		if slice[index].Quantity <= 0 {
			if len(slice) == 1 {
				gs.Characters[characterID].Inventory.Food = nil
			} else {
				slice[index] = slice[len(slice)-1]
				gs.Characters[characterID].Inventory.Food = slice[:len(slice)-1]
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
