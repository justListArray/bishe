package base

import "gorm.io/gorm"

type BaseController struct {
	db *gorm.DB
}
