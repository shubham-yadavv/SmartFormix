package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	ResponseID uint
	QuestionID uint
	Text       string
}
