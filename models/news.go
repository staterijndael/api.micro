package models

import (
	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Annotation string `gorm:"not null"`
	Body       string `gorm:"not null`
	Author_id  User   `gorm:"default:Anon"`
	Preview    string `gorm:"not null"`
	Background string `gorm:"default:null"`
	Types      string `gorm:"default:Системные"`
}

func (n *News) View() News {
	// return news with private settings
	return News{
		Title:      n.Title,
		Annotation: n.Annotation,
		Author_id:  n.Author_id,
		Preview:    n.Preview,
		Background: n.Background,
		Types:      n.Types,
	}
}
