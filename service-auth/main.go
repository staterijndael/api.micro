package main

import (
	"github.com/deissh/api.micro/models"
	"github.com/deissh/api.micro/service-auth/common"
	service "github.com/deissh/api.micro/service-auth/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"os"

	_ "github.com/deissh/api.micro/service-auth/docs"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// @title Service Auth API
// @version 1.0
// @description Auth, create tokens, and refresh old

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	conn := common.Init()

	// create tables if not exist
	// todo: add auto migration
	conn.AutoMigrate(&models.Token{}, &models.User{})

	handlers := service.CreateHandlers(conn)

	e.GET("/docs/*", echoSwagger.WrapHandler)
	g := e.Group("/api")
	{
		g.GET("/health", handlers.HealthCheckHandler)
		g.GET("/ping", handlers.PingHandler)

		g.POST("/auth", handlers.CreateHandler)
	}

	if err := e.Start(getEnv("HTTP_HOST", ":8080")); err != nil {
		e.Logger.Fatal(err)
	}
}
