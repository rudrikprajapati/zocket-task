package db

import (
	"task-manager-backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB(connectionString string) {
	var err error
	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	DB.AutoMigrate(&models.User{}, &models.Task{})
}
