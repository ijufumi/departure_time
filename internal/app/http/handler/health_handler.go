package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler interface {
	Health(ctx *gin.Context)
}

type healthHandler struct {
}

func (handler *healthHandler) Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}
