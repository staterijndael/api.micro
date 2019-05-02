package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RemoveRequest struct {
	// API version
	Version     string `json:"v" query:"v"`
	AccessToken string `form:"access_token" binding:"required"`
	All         bool   `form:"all"`
}

type RemoveResponse struct {
	// API version
	Version string `json:"v"`
	Status  string `json:"status"`
}

// TokenRemove godoc
// @Summary Remove token or all
// @Description Remove access_token or revome all tokens from this user
// @ID remove-token
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param access_token query string false "user access_token"
// @Param all query bool false "remove all tokens"
// @Success 200 {object} handlers.RemoveResponse
// @Success 400 {object} handlers.ResponseData
// @Router /token.remove [Get]
func (h Handler) RemoveHandler(c *gin.Context) {
	var r RemoveRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var token models.Token
	if err := h.db.Where(
		&models.Token{
			AccessToken: r.AccessToken,
		},
	).First(&token).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Access token not founded",
		})
		return
	}
	// remove token
	h.db.Delete(&token)

	if r.All == true {
		h.db.Where(&models.Token{
			UserId: token.UserId,
		}).Delete(models.Token{})
	}

	c.JSON(http.StatusOK, RemoveResponse{
		Version: "1",
		Status:  "ok",
	})
}
