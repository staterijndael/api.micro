package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	Version string `json:"v"`
	Ping    Ping   `json:"ping"`
}

// PingCheck godoc
// @Summary Ping service
// @ID ping-service
// @Accept  json
// @Produce  json
// @Success 200 {object} handlers.PingResponse
// @Router /_/ping [get]
func (h Handler) PingHandler(c *gin.Context) {
	ping := Ping{
		ID:          uuid.New().String(),
		ServiceName: "service-auth",
		Time:        time.Now().Local(),
	}

	c.JSON(http.StatusOK, PingResponse{
		Version: "1",
		Ping:    ping,
	})
}
