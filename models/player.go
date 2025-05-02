package models

type Player struct {
	Id         int    `json:"id" gorm:"int"`
	Name       string `json:"name" gorm:"varchar(255)"`
	Goal       int    `json:"goal" gorm:"int"`
	Pass       int    `json:"pass" gorm:"int"`
	Tackle     int    `json:"tackle" gorm:"int"`
	Foul       int    `json:"foul" gorm:"int"`
	Yellowcard int    `json:"yellowcard" gorm:"int"`
	Redcard    int    `json:"redcard" gorm:"int"`
}

func TableName() string {
	return "players"
}
