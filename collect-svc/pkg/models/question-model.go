package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	FormID uint
	Text   string
}
