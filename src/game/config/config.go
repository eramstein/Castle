package config

const WorldAgentID = -1

const (
	WhereInventory = iota
	WhereFloor     = iota
	WhereContainer = iota
)

const WorldSize = 5
const RegionWidth = 250
const RegionHeight = 250

const UpdateNeedsFrequency = 3600
const UpdateWeatherFrequency = 4 * 3600
