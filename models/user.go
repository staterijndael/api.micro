package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Nickname     string `gorm:"unique;not null;index:nickname"`
	Email        string `gorm:"unique;not null;index:email;type:varchar(100)"`
	Role         string `gorm:"not null;default:user"`
	Sex          int    `gorm:"not null;default:2"` // 1 – female; 2 – male.
	BDate        *time.Time
	Picture      string
	Desc         string
	Status       string
	Badges       []Badges
	PasswordHash string `gorm:"not null"`
}

type Badges struct {
	Name string `gorm:"not null"`
	Icon string `gorm:"not null"`
}

func (u *User) View() User {
	// return user with private settings
	return User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Nickname:  u.Nickname,
		Email:     u.Email,
		Role:      u.Role,
		Status:    u.Status,
		Badges:    u.Badges,
		Sex:       u.Sex,
		Picture:   u.Picture,
		Desc:      u.Desc,
		BDate:     u.BDate,
	}
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
