package models

type Member struct { //具体
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"type:varchar(255)"` //gorm:"uniqueIndex
	Identity string `json:"identity" gorm:"type:varchar(255)"` //(角色，如球员、教练、管理员等)
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Age      int    `json:"age" gorm:"type:int"`
	PlayerId int    `json:"player_id" gorm:"type:int"`
	// Position      string `json:"position" gorm:"type:varchar(255)"`
	// Jersey_number int    `json:"jersey_number" gorm:"type:int"`
}

func (Member) TableName() string {
	return "members"
}
