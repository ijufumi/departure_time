package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/email-service/internal/app/http/handler"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		healthHandler := handler.NewHealthHandler()
		api.GET("/health", healthHandler.Health)

		sendMailHandler := handler.NewSendMailHandler()
		api.POST("/send", sendMailHandler.Send)
	}
	_ = router.Run(":8080")
}
