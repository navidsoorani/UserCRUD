package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
