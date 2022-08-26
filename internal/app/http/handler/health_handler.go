package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandlerInterface interface {
	Health(ctx *gin.Context)
}

type healthHandler struct {
}

func (handler *healthHandler) Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func NewHealthHandler() HealthHandlerInterface {
	return &healthHandler{}
}
