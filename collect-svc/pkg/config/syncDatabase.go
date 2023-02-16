package config

import "github.com/shubham/collect-svc/pkg/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Form{})
	DB.AutoMigrate(&models.Question{})
	DB.AutoMigrate(&models.Response{})
	DB.AutoMigrate(&models.Answer{})
}
