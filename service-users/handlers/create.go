package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateRequestUser struct {
	FirstName    string          `form:"firstname" binding:"required"`
	LastName     string          `form:"lastname" binding:"required"`
	Nickname     string          `form:"nickname" binding:"required"` // unique
	Email        string          `form:"email" binding:"required"`    // unique
	Role         string          `form:"role" binding:"required"`
	Sex          int             `form:"sex"` // 1 – female; 2 – male.
	BDate        *time.Time      `form:"bdate"`
	Picture      string          `form:"picture"`
	Desc         string          `form:"desc"`
	Status       string          `form:"status"`
	Badges       []models.Badges `form:"badges"`
	PasswordHash string          `form:"phash" binding:"required"`
}

func (h Handler) createUser(c *gin.Context) {

	var r CreateRequestUser

	var user models.User

	if err := h.db.Where(&models.User{Nickname: r.Nickname}).First(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Nickname already registered",
		})
	}

}
