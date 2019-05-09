package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Token struct {
	gorm.Model
	AccessToken  string         `gorm:"unique;not null"`
	RefreshToken string         `gorm:"unique;not null"`
	UserId       int            `gorm:"not null"`
	Permissions  pq.StringArray `gorm:"not null;type:varchar(64)[]"`
}
