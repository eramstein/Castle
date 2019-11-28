package model

type Character struct {
	Pos            Pos
	Atts           Attributes
	Wounds         []Wound
	Diseases       []Disease
	Physical       PhysicalState
	Mental         MentalState
	Needs          NeedsState
	Desires        DesiresState
	NeedsProfile   NeedsProfile
	DesiresProfile DesiresProfile
	Task           CharacterTask
}

type Attributes struct {
	Strength     int
	Agility      int
	Dexterity    int
	Stamina      int
	Constitution int
	Immunity     int
	Reasonning   int
	Memory       int
	Focus        int
	Learning     int
	Creativity   int
}

// Needs
// if not satisfied, impact on physical state
type NeedsProfile struct {
	Hunger int
	Thirst int
	Sleep  int
	Warmth int
}

type NeedsState struct {
	Hunger int
	Thirst int
	Sleep  int
	Warmth int
}

// Desires
// if not satisfied, impact on mental state
type DesiresProfile struct {
}

type DesiresState struct {
}

type PhysicalState struct {
	Alive         bool
	Energy        int
	Focus         int
	Nutrition     NutritionState
	Hydratation   int
	Alcoolisation int
	Oxygenation   int
	Vitamins      VitaminsState
}

type NutritionState struct {
	Total    int
	Lipids   int
	Glucids  int
	Proteins int
}

type VitaminsState struct {
	C int
	D int
}

type Wound struct {
}

type Disease struct {
}

type MentalState struct {
}

type Ambitions struct {
}
