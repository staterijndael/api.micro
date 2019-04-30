package models

import (
	"time"
)

type Token struct {
	ID           string `gorm:"primary_key"`
	AccessToken  string `gorm:"unique;not null"`
	RefreshToken string `gorm:"unique;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
