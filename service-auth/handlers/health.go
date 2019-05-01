package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h Handler) HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseData{
		Status: http.StatusOK,
		Data: struct {
			Alive bool `json:"alive"`
		}{
			Alive: true,
		},
	})
}
