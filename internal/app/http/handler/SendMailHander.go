package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendMailHandlerInterface interface {
	Send(ctx *gin.Context)
}

type sendMailHandler struct {
}

func (handler *sendMailHandler) Send(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func NewSendMailHandler() SendMailHandlerInterface {
	return &sendMailHandler{}
}
