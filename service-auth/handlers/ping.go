package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func (h Handler) PingHandler(c echo.Context) error {
	ping := ResponsePing{
		ID:          uuid.New().String(),
		ServiceName: "service-auth",
		Time:        time.Now().Local(),
	}

	return c.JSON(http.StatusOK, ResponseData{
		Status: http.StatusOK,
		Data:   ping,
	})
}
