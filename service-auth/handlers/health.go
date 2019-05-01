package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Health struct {
	Alive bool `json:"alive"`
}

type HealthResponse struct {
	// API version
	Version string `json:"v"`
	Health  Health `json:"Health"`
}

// HealthCheck godoc
// @Summary Show health service
// @ID get-service-health
// @Accept  json
// @Produce  json
// @Success 200 {object} handlers.HealthResponse
// @Router /_/health [get]
func (h Handler) HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Version: "1",
		Health: Health{
			Alive: true,
		},
	})
}
