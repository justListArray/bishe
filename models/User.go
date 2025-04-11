// models/user.go
package models

type User struct {
	Username string `json:"username" gorm:"type:varchar(255)"` //gorm:"uniqueIndex
	Password string `json:"password" gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "user"
}
