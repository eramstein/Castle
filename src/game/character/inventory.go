package character

import (
	c "github.com/castle/src/game/config"
	"github.com/castle/src/game/item"
	m "github.com/castle/src/game/model"
)

func GetTimeToPickup() (seconds int) {
	return 10
}

func Pickup(gs *m.State, agentID int, quantity int, tile *m.Tile, characterID int, where int, index int, itemType int) {
	// update inventory
	switch itemType {
	case c.ItemTypeFood:
		var food m.Food
		switch where {
		case c.WhereFloor:
			food = tile.Items.Food[index]
			food.Quantity = quantity
		}
		item.AddFoodToStack(gs, food, nil, characterID, c.WhereInventory)
	}
	// update origin slice
	item.UpdateStackQuantity(gs, -quantity, tile, 0, where, index, itemType)
}
