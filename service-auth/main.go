package main

import (
	"github.com/deissh/api.micro/models"
	service "github.com/deissh/api.micro/service-auth/handlers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	//defer db.Close()

	return db, err
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	conn, err := ConnectDB()
	if err != nil {
		e.Logger.Panic(err)
	}

	// create tables if not exist
	// todo: add auto migration
	conn.AutoMigrate(&models.Token{}, &models.User{})

	handlers := service.CreateHandlers(conn)

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
