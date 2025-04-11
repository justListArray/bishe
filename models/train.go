package models

type Train struct {
	UserId    int    `json:"user_id" gorm:"int"`
	Name      string `json:"name" gorm:"type:varchar(255)"`
	Date      string `json:"date" gorm:"type:varchar(255)"`
	Content   string `json:"content" gorm:"type:varchar(255)"`
	Intensity string `json:"intensity" gorm:"type:varchar(255)"`
	Duration  int    `json:"duration" gorm:"int"`
	Injury    bool   `json:"injury" gorm:"bool"`
}

func (Train) TableName() string {
	return "Train"
}
