package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/departure_time/internal/app/http/handler"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	api := router.Group("/api/v1")
	{
		healthHandler := handler.NewHealthHandler()
		api.GET("/health", healthHandler.Health)
	}
	_ = router.Run(":8080")
}
