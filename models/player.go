package models

type Player struct {
	id         int
	name       string
	Goal       int
	Pass       int
	Tackle     int
	Foul       int
	Yellowcard int
	Redcard    int
}

func TableName() string {
	return "PlayerData"
}
