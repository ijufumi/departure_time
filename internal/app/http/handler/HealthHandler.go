package handler

import "github.com/gin-gonic/gin"

type HealthHandlerInterface interface {
	Health(ctx *gin.Context)
}

type HealthHandler struct {
}

func (handler *HealthHandler) Health(ctx *gin.Context) {
}

func NewHealthHandler() HealthHandlerInterface {
	return &HealthHandler{}
}
