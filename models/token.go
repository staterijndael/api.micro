package models

import (
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	AccessToken  string `gorm:"unique;not null"`
	RefreshToken string `gorm:"unique;not null"`
	UserId       int    `gorm:"not null"`
}
