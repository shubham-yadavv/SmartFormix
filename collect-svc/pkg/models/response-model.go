package models

import "gorm.io/gorm"

type Response struct {
	gorm.Model
	FormID uint
}
