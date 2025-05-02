package models

type Train struct {
	TrainId   int    `json:"trainId" gorm:"int"`
	UserId    int    `json:"userId" gorm:"int"`
	Name      string `json:"name" gorm:"type:varchar(255)"`
	Date      string `json:"date" gorm:"type:varchar(255)"` // dateLayout := "2006-01-02" // *定义日期格式*
	Content   string `json:"content" gorm:"type:varchar(255)"`
	Intensity string `json:"intensity" gorm:"type:varchar(255)"`
	Duration  int    `json:"duration" gorm:"int"`
	Injury    bool   `json:"injury" gorm:"bool"`
}

func (Train) TableName() string {
	return "Train"
}
