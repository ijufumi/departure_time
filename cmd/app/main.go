package app

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/email-service/internal/app/container"
)

func RunApp() error {
	router := gin.Default()

	c := container.NewContainer()

	router.Use(static.Serve("/", static.LocalFile("./web/public", true)))

	api := router.Group("/api/v1")
	{
		api.GET("/health", c.GetHealthHandler().Health)
		api.POST("/mail/send", c.GetSendMailHandler().Send)
	}
	return router.Run(":8080")
}
