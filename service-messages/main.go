package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-messages/common"
	service "github.com/deissh/api.micro/service-messages/handlers"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/deissh/api.micro/service-messages/docs"
)

// @title Service Messages API
// @version 1.0
// @description Messages methods

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.GET("/messages.addChatUser")
		g.GET("/messages.createChat")
		g.GET("/messages.delete")
		g.GET("/messages.deleteChatPhoto")
		g.GET("/messages.setChatPhoto")
		g.GET("/messages.deleteConversation")
		g.GET("/messages.edit")
		g.GET("/messages.editChat")
		g.GET("/messages.getById")
		g.GET("/messages.getChat")
		g.GET("/messages.getConversationMembers")
		g.GET("/messages.getConversations")
		g.GET("/messages.getConversationsById")
		g.GET("/messages.getHistory")
		g.GET("/messages.getLastActivity")
		g.GET("/messages.getInviteLink")
		g.GET("/messages.joinChatByInviteLink")
		g.GET("/messages.markAsRead")
		g.GET("/messages.removeChatUser")
		g.GET("/messages.send")

		g.GET("/_/health", handlers.HealthCheckHandler)
		g.GET("/_/ping", handlers.PingHandler)
		g.GET("/_/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
