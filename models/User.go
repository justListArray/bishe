// models/user.go
package models

type User struct {
	Id       int    `json:"id" gorm:"type:int"`
	Username string `json:"username" gorm:"type:varchar(255)"` //gorm:"uniqueIndex
	Password string `json:"password" gorm:"type:varchar(255)"`
	Identity string `json:"identity" gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "users"
}
