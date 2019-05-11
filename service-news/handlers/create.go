package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateNewsR struct {
	Title      string `form:"title" binding:"required"`
	Annotation string `form:"annotation" binding:"required"`
	Body       string `form:"body" binding:"required"`
	Author_id  string `form:"author_id"`
	Preview    string `form:"preview" binding:"required"`
	Background string `form:"background"`
	Types      string `form:"types"`
}

func (h Handler) CreateNews(c *gin.Context) {
	var r CreateNewsR
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	type CreateResponse struct {
		// API version
		Version string      `json:"v"`
		News    models.News `json:"news"`
	}

	var author models.User

	if r.Author_id != "" {
		err := h.db.Where(&models.User{Nickname: r.Author_id}).First(&author)
		if err != nil {
			c.JSON(http.StatusBadRequest, ResponseData{
				Status: http.StatusBadRequest,
				Data:   "Bad user nickname",
			})
			return
		}
	} else {

		author = models.User{

			FirstName:    "anon",
			LastName:     "anon",
			Nickname:     "anon",
			Email:        "anon",
			Role:         "anon",
			Sex:          1,
			Picture:      "anon",
			Desc:         "anon",
			Status:       "anon",
			PasswordHash: "anon",
		}
	}

	news := models.News{
		Title:      r.Title,
		Annotation: r.Annotation,
		Body:       r.Body,
		Author_id:  author,
		Preview:    r.Preview,
		Background: r.Background,
		Types:      r.Types,
	}

	h.db.Create(&news)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		News:    news,
	})
}
