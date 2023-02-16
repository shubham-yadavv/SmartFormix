package models

import "gorm.io/gorm"

type Form struct {
	gorm.Model
	Name        string
	Description string
}
