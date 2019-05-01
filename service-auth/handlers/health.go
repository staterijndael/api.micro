package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
	Alive bool `json:"alive"`
}

type HealthResponse struct {
	// API version
	Api    string `json:"api"`
	Health Health `json:"Health"`
}

// HealthCheck godoc
// @Summary Show health service
// @ID get-service-health
// @Accept  json
// @Produce  json
// @Success 200 {object} handlers.HealthResponse
// @Router /api/health [get]
func (h Handler) HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{
		Api: "v1",
		Health: Health{
			Alive: true,
		},
	})
}
