package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type Task struct {
	gorm.Model
	Title       string
	Description string
	AssignedTo  uint
	Status      string
	CreatedBy   uint
}
