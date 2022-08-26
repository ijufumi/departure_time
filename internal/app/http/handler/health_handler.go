package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthHandler is the handler for health check from load balancer
type HealthHandler interface {
	Health(ctx *gin.Context)
}

type healthHandler struct {
}

// Health returns always 200
func (handler *healthHandler) Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

// NewHealthHandler is factory method of HealthHandler
func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}
