package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RemoveRequest struct {
	// API version
	Version     string `json:"v" query:"v"`
	AccessToken string `query:"access_token" binding:"required"`
	All         bool   `query:"all" binding:"required"`
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
// @Router /token.remove [Get]
func (h Handler) RemoveHandler(c *gin.Context) {
	c.JSON(http.StatusOK, RemoveResponse{
		Version: "1",
		Status:  "ok",
	})
}
