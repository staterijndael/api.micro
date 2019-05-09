package handlers

import (
	"github.com/Oringik/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateRequest struct {
	Q     string `form:"q"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}

func (h Handler) getNews(c *gin.Context) {

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var news models.News
	if err := h.db.Where(&models.News{Title: r.Q}).First(&news).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad title name",
		})
		return
	}

}
