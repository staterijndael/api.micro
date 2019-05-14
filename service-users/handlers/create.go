package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type CreateRequestUser struct {
	FirstName    string `form:"firstname" binding:"required"`
	LastName     string `form:"lastname" binding:"required"`
	Nickname     string `form:"nickname" binding:"required"` // unique
	Email        string `form:"email" binding:"required"`    // unique
	Role         string `form:"role" binding:"required"`
	Sex          int    `form:"sex"` // 1 – female; 2 – male.
	BDate        string `form:"bdate"`
	Picture      string `form:"picture"`
	Desc         string `form:"desc"`
	Status       string `form:"status"`
	Badges       string `form:"badges"`
	PasswordHash string `form:"phash" binding:"required"`
}

type uCreateResponse struct {
	Version      string      `json:"v"`
	User         models.User `json:"user"`
	FailedBadges []string    `json:"failedbadges"`
}

func (h Handler) CreateUser(c *gin.Context) {

	var r CreateRequestUser

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var user models.User

	if err := h.db.Where(&models.User{Nickname: r.Nickname}).First(&user).Error; err == nil {

		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Nickname already registered",
		})
		return
	}

	if err := h.db.Where(&models.User{Email: r.Email}).First(&user).Error; err == nil {

		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Email already registered",
		})
		return
	}

	Badges := strings.Split(r.Badges, ",")

	var b models.Badges

	var b_array []models.Badges

	var failedBadges []string

	for _, k := range Badges {
		if err := h.db.Where(&models.Badges{Name: k}).First(&b).Error; err == nil {
			b_array = append(b_array, b)
		} else {
			failedBadges = append(failedBadges, k)
		}
	}

	d, _ := time.Parse("2006-01-02", r.BDate)

	us := models.User{
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		Nickname:     r.Nickname,
		Email:        r.Email,
		Role:         r.Role,
		Sex:          r.Sex,
		BDate:        d,
		Picture:      r.Picture,
		Desc:         r.Desc,
		Status:       r.Status,
		Badges:       b_array,
		PasswordHash: r.PasswordHash,
	}

	h.db.Create(&us)

	c.JSON(http.StatusOK, uCreateResponse{
		Version:      "1",
		User:         us,
		FailedBadges: failedBadges,
	})
}
