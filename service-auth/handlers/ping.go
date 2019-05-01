package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Ping struct {
	ID          string    `json:"id,omitempty"`
	ServiceName string    `json:"service,omitempty"`
	Time        time.Time `json:"time,omitempty"`
}

type PingResponse struct {
	// API version
	Api  string `json:"api"`
	Ping Ping   `json:"ping"`
}

// PingCheck godoc
// @Summary Ping service
// @ID ping-service
// @Accept  json
// @Produce  json
// @Success 200 {object} handlers.PingResponse
// @Router /api/ping [get]
func (h Handler) PingHandler(c echo.Context) error {
	ping := Ping{
		ID:          uuid.New().String(),
		ServiceName: "service-auth",
		Time:        time.Now().Local(),
	}

	return c.JSON(http.StatusOK, PingResponse{
		Api:  "v1",
		Ping: ping,
	})
}
