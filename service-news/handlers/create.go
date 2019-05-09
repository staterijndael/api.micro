package handlers

import (
	"github.com/Oringik/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateNewsR struct {
	title      string `form:"title" binding:"required"`
	annotation string `form:"title" binding:"required"`
	body       string `gorm:"form:"body" required:"required"`
	author_id  string `form:"author_id"`
	preview    string `form:"preview" binding:"required"`
	background string `form:"background"`
	types      string `form:"types"`
}

func (h Handler) createNews(c *gin.Context) {
	var r CreateNewsR
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var n CreateNewsR
	var author models.User

	err := h.db.Where(&models.User{Nickname: r.author_id}).First(&author)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad user nickname",
		})
		return
	}

	news := models.News{
		Title:      n.Title,
		Annotation: n.Annotation,
		Body:       n.Body,
		Author_id:  Author,
		Preview:    n.Preview,
		Background: n.Background,
		Types:      n.Types,
	}

	h.db.Create(&news)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		News:    news,
	})
}
