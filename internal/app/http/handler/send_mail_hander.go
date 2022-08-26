package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/email-service/internal/app/http/request"
	"github.com/ijufumi/email-service/internal/app/service"
	"net/http"
)

// SendMailHandler is the handler for sending mail
type SendMailHandler interface {
	Send(ctx *gin.Context)
}

type sendMailHandler struct {
	sendMailService service.SendMailService
}

// Send allows sending email via Email Vendors
func (handler *sendMailHandler) Send(ctx *gin.Context) {
	var contents request.SendMail
	if err := ctx.BindJSON(&contents); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := handler.sendMailService.Send(contents); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusOK)
}

// NewSendMailHandler is factory method of SendMailHandler
func NewSendMailHandler(sendMailService service.SendMailService) SendMailHandler {
	return &sendMailHandler{sendMailService: sendMailService}
}
