package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefRequest struct {
	// API version
	Version    string `form:"v"`
	Title      string `form:"title"`
	Annotation string `form:"annotation"`
	Body       string `form:"body"`
	Author_id  string `form:"author_id"`
	Preview    string `form:"preview"`
	Background string `form:"background"`
	Types      string `form:"types"`
}

type RefResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

func checkNull(old string, new string) string {
	if new == "" {
		new = old
	}
	return new
}

func (h Handler) RefreshNews(c *gin.Context) {
	var r RefRequest
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

	var author models.User

	err := h.db.Where(&models.User{Nickname: r.Author_id}).First(&author)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad user nickname",
		})
		return
	}

	newNews := models.News{
		Title:      checkNull(news.Title, r.Title),
		Annotation: checkNull(news.Annotation, r.Annotation),
		Body:       checkNull(news.Body, r.Body),
		Author_id:  author,
		Preview:    checkNull(news.Preview, r.Preview),
		Background: checkNull(news.Background, r.Background),
		Types:      checkNull(news.Types, r.Types),
	}

	h.db.Create(&newNews)

	c.JSON(http.StatusOK, RefResponse{
		Version: "1",
		News:    newNews,
	})
}
