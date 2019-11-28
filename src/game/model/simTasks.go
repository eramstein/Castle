package model

type SimTaskSchedule struct {
	EndTime int
	Type    int
	Agent   int
}

type CharacterTask struct {
	LastUpdated int
	Completion  float32
	Type        int
	SubType     int
	//data
	Agent int
	Pos   Pos
	Zone  int
	Item  int
	// ...
}

const (
	TaskTypeUpdateNeeds   = 1
	TaskTypeUpdateWeather = 2
	TaskTypeMovement      = 3
)
