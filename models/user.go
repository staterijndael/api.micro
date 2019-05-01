package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName       string `gorm:"not null"`
	LastName        string `gorm:"not null"`
	Nickname        string `gorm:"unique;not null;index:nickname"`
	Email           string `gorm:"unique;not null;index:email;type:varchar(100)"`
	Sex             int    `gorm:"not null;default:2"` // 1 – female; 2 – male.
	BDate           *time.Time
	BDateVisibility int `gorm:"default:2"` // 1 – show birth date; 2 – show only month and day; 0 – hide birth date.
	Picture         string
	Desc            string
	Status          string
	Badges          []Badges
	Password        string `gorm:"not null"`
}

type Badges struct {
	Name string `gorm:"not null"`
	Icon string `gorm:"not null"`
}
