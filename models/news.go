package models

import (
	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	Title      string `gorm:"not null" json:"title"`
	Annotation string `gorm:"not null" json:"annotation"`
	Body       string `gorm:"not null	json:"body"`
	Author_id  User   `default:"Anon"	json:"author_id"`
	Preview    string `gorm:"not null" json:"preview"`
	Background string `default:"null" json:"background"`
	Types      string `default:"Системные" json:"types"`
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
