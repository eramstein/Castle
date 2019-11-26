package player

import (
	m "github.com/castle/src/game/model"
)

func GetInitialPlayer() *m.Player {
	return &m.Player{
		Pos: m.Pos{
			X:      15,
			Y:      15,
			Z:      0,
			Region: 0,
		},
		Atts: m.Attributes{
			Strength:     0,
			Agility:      0,
			Dexterity:    0,
			Stamina:      0,
			Constitution: 0,
			Immunity:     0,
			Reasonning:   0,
			Memory:       0,
			Focus:        0,
			Learning:     0,
			Creativity:   0,
		},
		Wounds:   nil,
		Diseases: nil,
		Physical: m.PhysicalState{
			Alive:  false,
			Energy: 0,
			Focus:  0,
			Nutrition: m.NutritionState{
				Total:    0,
				Lipids:   0,
				Glucids:  0,
				Proteins: 0,
			},
			Hydratation:   0,
			Alcoolisation: 0,
			Oxygenation:   0,
			Vitamins: m.VitaminsState{
				C: 0,
				D: 0,
			},
		},
		Needs: m.NeedsState{
			Hunger: 0,
			Thirst: 0,
			Sleep:  0,
			Warmth: 0,
		},
		NeedsProfile: m.NeedsProfile{
			Hunger: 100,
			Thirst: 100,
			Sleep:  100,
			Warmth: 100,
		},
	}
}
