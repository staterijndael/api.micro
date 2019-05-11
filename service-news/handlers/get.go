package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateRequest struct {
	Q     string `form:"q"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}

type CreateResponse struct {
	// API version
	Version string        `json:"v"`
	News    []models.News `json:"news"`
}

func (h Handler) GetNews(c *gin.Context) {

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	news := []models.News{}
	if r.Q == "" {
		h.db.Find(&news)
	} else {
		if err := h.db.Where(&models.News{Title: r.Q}).First(&news).Error; err != nil {
			c.JSON(http.StatusBadRequest, ResponseData{
				Status: http.StatusBadRequest,
				Data:   "Bad title name",
			})
			return
		}
	}

	if r.Limit != 0 && r.Page != 0 {
		if r.Limit*r.Page > len(news) {
			c.JSON(http.StatusBadRequest, ResponseData{
				Status: http.StatusBadRequest,
				Data:   "Array index error",
			})
			return
		} else {
			news = news[r.Limit*r.Page-r.Limit : r.Limit*r.Page]
		}
	}

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		News:    news,
	})
}
