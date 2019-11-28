package model

type Player struct {
	Pos          Pos
	Atts         Attributes
	Wounds       []Wound
	Diseases     []Disease
	Physical     PhysicalState
	Needs        NeedsState
	NeedsProfile NeedsProfile
	Task         CharacterTask
}
