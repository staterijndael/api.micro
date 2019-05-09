package handlers

import (
	"github.com/Oringik/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshRequest struct {
	// API version
	Version    string `form:"v"`
	Title      string `form:"title" binding:"required"`
	Annotation string `form:"title" binding:"required"`
	Body       string `gorm:"form:"body" required:"required"`
	Author_id  string `form:"author_id"`
	Preview    string `form:"preview" binding:"required"`
	Background string `form:"background"`
	Types      string `form:"types"`
}

type RefreshResponse struct {
	// API version
	Version string       `json:"v"`
	Token   models.Token `json:"token"`
}

func (h Handler) RefreshNews(c *gin.Context) {
	var r RefreshRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var news models.News
	if err := h.db.Where(
		&models.News{
			Title: r.Title,
		},
	).First(&news).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "News did not find",
		})
		return
	}

	h.db.Delete(&news)

	var author models.User

	err := h.db.Where(&models.User{Nickname: r.Author_id}).First(&author)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad user nickname",
		})
		return
	}

	var n RefreshRequest

	newNews := models.News{
		Title:      n.Title,
		Annotation: n.Annotation,
		Body:       n.Body,
		Author_id:  author,
		Preview:    n.Preview,
		Background: n.Background,
		Types:      n.Types,
	}

	h.db.Create(&newNews)

	c.JSON(http.StatusOK, RefreshResponse{
		Version: "1",
		Token:   newNews,
	})
}
