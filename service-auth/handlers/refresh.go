package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshRequest struct {
	// API version
	Version      string `form:"v"`
	RefreshToken string `form:"refresh_token" binding:"required"`
}

type RefreshResponse struct {
	// API version
	Version string       `json:"v"`
	Token   models.Token `json:"token"`
}

// TokenRefresh godoc
// @Summary Deactivate old token and create new
// @Description Generate new access_token and refresh_token
// @ID refresh-token
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param refresh_token query string false "refresh_token"
// @Success 200 {object} handlers.RefreshResponse
// @Failure 400 {object} handlers.ResponseData
// @Failure 500 {object} handlers.ResponseData
// @Router /token.refresg [Get]
func (h Handler) RefreshHandler(c *gin.Context) {
	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		Token:   models.Token{},
	})
}
