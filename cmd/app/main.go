package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/email-service/internal/app/container"
)

func main() {
	router := gin.Default()

	c := container.NewContainer()

	router.Use(static.Serve("/", static.LocalFile("./web/public", true)))

	api := router.Group("/api/v1")
	{
		api.GET("/health", c.GetHealthHandler().Health)
		api.POST("/mail/send", c.GetSendMailHandler().Send)
	}
	_ = router.Run(":8080")
}
