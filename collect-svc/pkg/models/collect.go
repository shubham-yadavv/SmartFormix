package models

import "time"

type Form struct {
	FormID    uint       `gorm:"primaryKey;autoIncrement" json:"form_id"`
	FormName  string     `json:"form_name"`
	FormDesc  string     `json:"form_desc"`
	Questions []Question `gorm:"foreignKey:FormID" json:"questions"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Question struct {
	QuestionID   uint      `gorm:"primaryKey;autoIncrement" json:"question_id"`
	FormID       uint      `json:"form_id"`
	QuestionText string    `json:"question_text"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Response struct {
	ResponseID     uint      `gorm:"primaryKey;autoIncrement" json:"response_id"`
	FormID         uint      `json:"form_id"`
	Timestamp      time.Time `json:"timestamp"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Answers        []Answer  `gorm:"foreignKey:ResponseID" json:"answers"`
}
type Answer struct {
	AnswerID   uint      `gorm:"primaryKey;autoIncrement" json:"answer_id"`
	ResponseID uint      `json:"response_id"`
	QuestionID uint      `json:"question_id"`
	AnswerText string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
