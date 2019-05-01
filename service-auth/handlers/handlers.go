package handlers

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponsePing struct {
	ID          string    `json:"id,omitempty"`
	ServiceName string    `json:"service,omitempty"`
	Time        time.Time `json:"time,omitempty"`
}

type Handler struct {
	db *gorm.DB
}

func CreateHandlers(db *gorm.DB) Handler {
	return Handler{db}
}
