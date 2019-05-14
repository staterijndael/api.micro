package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RemoveRequest struct {
	Q string `form:"q"`
}

type RemoveResponse struct {
	User models.User `json:"user"`
}

func (h Handler) RemoveNews(c *gin.Context) {
	var r RemoveRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var user models.User
	if err := h.db.Where(
		&models.User{
			Nickname: r.Q,
		},
	).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "User not founded",
		})
		return
	}

	h.db.Delete(&user)

	c.JSON(http.StatusOK, RemoveResponse{
		User: user,
	})
}
