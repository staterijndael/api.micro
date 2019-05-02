package models

import (
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	AccessToken  string           `gorm:"unique;not null"`
	RefreshToken string           `gorm:"unique;not null"`
	UserId       int              `gorm:"not null"`
	Permissions  TokenPermissions `gorm:"not null"`
}

type TokenPermissions struct {
	Notify        bool `gorm:"default:false"`
	Friends       bool `gorm:"default:false"`
	Status        bool `gorm:"default:false"`
	Messages      bool `gorm:"default:false"`
	Comments      bool `gorm:"default:false"`
	Wall          bool `gorm:"default:false"`
	Notifications bool `gorm:"default:false"`
	Email         bool `gorm:"default:false"`
}

func (t Token) AddScopes(scope []string) {

}
