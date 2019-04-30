package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Token struct {
	gorm.Model
	ID           string `gorm:"primary_key"`
	AccessToken  string `gorm:"unique;not null"`
	RefreshToken string `gorm:"unique;not null"`
	UserId 		 string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
