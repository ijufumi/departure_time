package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/email-service/internal/app/container"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	c := container.NewContainer()

	api := router.Group("/api/v1")
	{
		api.GET("/health", c.GetHealthHandler().Health)
		api.POST("/send", c.GetSendMailHandler().Send)
	}
	_ = router.Run(":8080")
}
