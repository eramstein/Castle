package config

const PlayerAgentID = 0
const WorldAgentID = -1

const (
	WhereInventory = iota
	WhereFloor     = iota
	WhereContainer = iota
)

const (
	ItemTypeFood = iota
)

const (
	AgentTypeCharacter = iota
	AgentTypeAnimal    = iota
	AgentTypePhantom   = iota
)

const WorldSize = 5
const RegionWidth = 250
const RegionHeight = 250

const UpdateNeedsFrequency = 3600
const UpdateWeatherFrequency = 4 * 3600
