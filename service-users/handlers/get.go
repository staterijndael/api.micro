package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type gCreateRequest struct {
	Q string `form:"q"`
}

type gCreateResponse struct {
	Users []models.User `json:"users"`
}

func (h Handler) GetUser(c *gin.Context) {

	var r gCreateRequest
	var user []models.User

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	if r.Q == "" {
		h.db.Find(&user)
	} else {
		if err := h.db.Where(&models.User{Nickname: r.Q}).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, ResponseData{
				Status: http.StatusBadRequest,
				Data:   "Bad nickname",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gCreateResponse{
		Users: user,
	})
}
