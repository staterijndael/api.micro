package main

import (
	"github.com/deissh/api.micro/service-auth/common"
	service "github.com/deissh/api.micro/service-auth/handlers"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	r := gin.Default()

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	r.GET("/token.create", handlers.CreateHandler)
	r.GET("/token.refresh", handlers.CreateHandler)
	r.GET("/token.remove", handlers.CreateHandler)

	g := r.Group("/_")
	{
		// additional methods
		g.GET("/health", handlers.HealthCheckHandler)
		g.GET("/ping", handlers.PingHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.Use(gin.Recovery())

	if err := r.Run(getEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
