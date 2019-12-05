package player

import (
	m "github.com/castle/src/game/model"
)

func GetInitialPlayerCharacter() *m.Character {
	return &m.Character{
		LivingBeing: m.LivingBeing{
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
				Alive:  true,
				Energy: 100,
				Focus:  100,
				Nutrition: m.NutritionState{
					Total:    100,
					Lipids:   100,
					Glucids:  100,
					Proteins: 100,
				},
				Hydratation:   100,
				Alcoolisation: 0,
				Oxygenation:   100,
				Vitamins: m.VitaminsState{
					C: 100,
					D: 100,
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
		},
	}
}
