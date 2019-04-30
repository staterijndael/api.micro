package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"time"
)

type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponsePing struct {
	ID          string    `json:"id,omitempty"`
	ServiceName string    `json:"service,omitempty"`
	Time        time.Time `json:"time,omitempty"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func PingHandler(c echo.Context) error {
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

func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseData{
		Status: http.StatusOK,
		Data: struct {
			Alive bool `json:"alive"`
		}{
			Alive: true,
		},
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	g := e.Group("/api")
	{
		g.GET("/health", HealthCheckHandler)
		g.GET("/ping", PingHandler)
	}

	if err := e.Start(getEnv("HTTP_HOST", ":8080")); err != nil {
		e.Logger.Fatal(err)
	}
}
