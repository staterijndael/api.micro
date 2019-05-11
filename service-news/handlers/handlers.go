package handlers

import (
	"github.com/jinzhu/gorm"
)

// Error default response type
type ResponseData struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
	Author string
}

type Handler struct {
	db *gorm.DB
}

func CreateHandlers(db *gorm.DB) Handler {
	return Handler{db}
}
